package main

import (
	"context"
	"fmt"
	"github.com/kechdarho/FinTrack/auth/internal/handlers"
	"github.com/kechdarho/FinTrack/auth/internal/services/auth"
	"github.com/kechdarho/FinTrack/auth/internal/storage/authPg"
	"github.com/kechdarho/FinTrack/auth/pkg/config"
	"github.com/kechdarho/FinTrack/auth/pkg/jwt"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		fmt.Println(config.Config)
		panic(err)
	}

	memoryCache := cache.New(
		config.Config.Cache.Memory.DefaultExpiration,
		config.Config.Cache.Memory.CleanupInterval,
	)

	authStorage, err := authPg.NewAuthStorage(
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.DbName,
		config.Config.Database.Username,
		config.Config.Database.Password,
		config.Config.Database.Sslmode,
	)

	if err != nil {
		panic(err)
	}

	defer authStorage.Close()
	jwtWrapper := jwt.NewJWTWrapper(config.Config.JWT.Secret, config.Config.JWT.AccessTokenTTL, config.Config.JWT.RefreshTokenTTL)
	authService := auth.NewAuthService(memoryCache, authStorage, jwtWrapper)

	router := handlers.NewHandlers(authService).Router

	srv := &http.Server{
		Addr:    config.Config.Server.Host + ":" + config.Config.Server.Port,
		Handler: router,
	}

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		log.Printf("Сервер запущен на %s", srv.Addr)
		if err = srv.ListenAndServe(); err != nil {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	<-quit

	log.Println("Получен сигнал завершения, остановка сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении сервера: %v", err)
	}

	log.Println("Сервер успешно остановлен")
}
