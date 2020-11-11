package httpserver

import (
	"net/http"

	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/storage"
)

type APIHandler struct {
	strg  storage.Storager
	event *EventHandler
}

func newAPIHandler(strg storage.Storager) *APIHandler {
	return &APIHandler{
		strg:  strg,
		event: NewEventHandler(strg),
	}
}

var _ http.Handler = (*APIHandler)(nil)

func (h *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)

	switch head {
	case "event":
		h.event.ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}
