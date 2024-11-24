package main

import (
	"github.com/kechdarho/FinTrack/auth/internal/handlers"
	"github.com/kechdarho/FinTrack/auth/internal/services/auth"
	"github.com/kechdarho/FinTrack/auth/internal/storage"
	"github.com/kechdarho/FinTrack/auth/pkg/config"
	"github.com/patrickmn/go-cache"
	"log"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		panic(err)
	}

	memoryCache := cache.New(
		config.Config.Cache.Memory.DefaultExpiration,
		config.Config.Cache.Memory.CleanupInterval,
	)

	authStorage, err := storage.NewAuthStorage(
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

	authService := auth.NewAuthService(memoryCache, authStorage)

	h := handlers.NewHandlers(authService)

	go func() {
		log.Printf("Сервер запущен на %s:%d", config.Config.Server.Host, config.Config.Server.Port)
		log.Fatal(h.Router.Run(config.Config.Server.Host + ":" + config.Config.Server.Port))

	}()

}
