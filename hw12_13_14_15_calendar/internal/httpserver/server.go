package httpserver

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/config"
)

func NewServerByConfig(cfg *config.Config) *http.Server {
	handler := &RootHandler{}
	return &http.Server{
		Addr:         cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Handler:      middlewareLogger(handler),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
