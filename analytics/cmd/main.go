package main

import (
	"github.com/kechdarho/FinTrack/auth/pkg/config"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		panic("loading config failed")
	}
	authService := auth.
}
