package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/config"
	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/logger"
)

func main() {
	wg := sync.WaitGroup{}

	cfg := config.Configure()
	cfg.ConfigureLoggerByConfig()
	cfg.RunStorage()

	quitCh := make(chan interface{})

	cfg.RunHTTPServer(quitCh, &wg)
	if err := cfg.RunGRPCServer(quitCh, &wg); err != nil {
		logger.Error(err)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-sigc
	close(quitCh)

	wg.Wait()
}
