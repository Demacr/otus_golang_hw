package main

import (
	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/config"
	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/httpserver"
	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/storage"
)

func main() {
	cfg := config.Configure()
	logger.ConfigureLoggerByConfig(cfg)

	storage.GetStoragerByConfig(cfg)
	defer logger.Close()

	server := httpserver.NewServerByConfig(cfg)
	err := server.ListenAndServe()
	if err != nil {
		logger.Error(err)
	}
}
