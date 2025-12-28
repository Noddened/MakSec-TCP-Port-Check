package main

import (
	"os"

	"github.com/Noddened/MakSec-TCP-Port-Check/internal/app"
	"github.com/Noddened/MakSec-TCP-Port-Check/internal/config"
	"github.com/Noddened/MakSec-TCP-Port-Check/internal/logger"
)

func main() {
	logger := logger.SetupLogger()

	cfg, err := config.LoadConfig(logger)
	if err != nil {
		logger.Error("Ошибка загрузки конфига:", err)
		os.Exit(1)
	}

	server := app.NewServer(cfg)
	if err := server.Start(logger); err != nil {
		logger.Error("Ошибка загрузки сервера:", err)
		os.Exit(1)
	}
}
