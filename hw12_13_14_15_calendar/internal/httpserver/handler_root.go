package httpserver

import (
	"net/http"

	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/storage"
)

type RootHandler struct {
	apiHandler *APIHandler
}

func newRootHandler(strg storage.Storager) *RootHandler {
	return &RootHandler{
		apiHandler: newAPIHandler(strg),
	}
}

var _ http.Handler = (*RootHandler)(nil)

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)

	switch head {
	case "":
		w.WriteHeader(http.StatusOK)
	case "api":
		h.apiHandler.ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}
