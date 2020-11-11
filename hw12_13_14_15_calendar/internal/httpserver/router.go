package httpserver

import (
	"net/http"

	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/storage"
)

type Router struct {
	rootHandler *RootHandler
}

func NewRouter(strg storage.Storager) *Router {
	return &Router{
		rootHandler: newRootHandler(strg),
	}
}

func (r *Router) RootHandler() http.Handler {
	return r.rootHandler
}
