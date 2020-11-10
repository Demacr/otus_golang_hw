package httpserver

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/storage"
)

type EventHandler struct {
	strg storage.Storager
}

func NewEventHandler(strg storage.Storager) *EventHandler {
	return &EventHandler{
		strg: strg,
	}
}

func (h *EventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)

	switch r.Method {
	case "GET":
		eventGet(h, w, r.URL)
	case "POST":
		eventPost(h, w, r.Body)
	case "PATCH":
		eventPatch(h, w, r.Body, head)
	case "DELETE":
		eventDelete(h, w, head)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func eventGet(h *EventHandler, w http.ResponseWriter, u *url.URL) {
	var ok bool
	var typeQ []string
	var datetimeQ []string
	if typeQ, ok = u.Query()["type"]; !ok || len(typeQ) < 1 {
		http.Error(w, "Missing type query", http.StatusBadRequest)
		return
	}
	if datetimeQ, ok = u.Query()["dt"]; !ok || len(datetimeQ) < 1 {
		http.Error(w, "Missing dt (datetime) query", http.StatusBadRequest)
	}
	dt, err := time.Parse("2006-01-02", datetimeQ[0])
	if err != nil {
		http.Error(w, "Wrong datetime format", http.StatusBadRequest)
		return
	}

	var events []storage.Event
	switch typeQ[0] {
	case "day":
		events = h.strg.ListDay(dt)
	case "week":
		events = h.strg.ListWeek(dt)
	case "month":
		events = h.strg.ListMonth(dt)
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	writeJSON(w, events)
}

func eventPost(h *EventHandler, w http.ResponseWriter, body io.ReadCloser) {
	defer body.Close()
	ar := AddRequest{}

	jd := json.NewDecoder(body)
	if err := jd.Decode(&ar); err != nil {
		http.Error(w, "wrong json", http.StatusInternalServerError)
		return
	}
	if err := h.strg.Add(&storage.Event{
		UUID:               ar.UUID,
		Header:             ar.Header,
		DateTime:           ar.DateTime,
		Duration:           storage.TDuration(ar.Duration),
		Description:        ar.Description,
		UserID:             ar.UserID,
		NotificationBefore: storage.TDuration(ar.NotificationBefore),
	}); err != nil {
		http.Error(w, "Error during addition event", http.StatusInternalServerError)
		return
	}
}

func eventPatch(h *EventHandler, w http.ResponseWriter, body io.ReadCloser, uuid string) {
	defer body.Close()
	mr := ModifyRequest{}

	jd := json.NewDecoder(body)
	if err := jd.Decode(&mr); err != nil {
		http.Error(w, "wrong json", http.StatusInternalServerError)
		return
	}
	if err := h.strg.Modify(uuid, &storage.Event{
		UUID:               uuid,
		Header:             mr.Header,
		DateTime:           mr.DateTime,
		Duration:           storage.TDuration(mr.Duration),
		Description:        mr.Description,
		UserID:             mr.UserID,
		NotificationBefore: storage.TDuration(mr.NotificationBefore),
	}); err != nil {
		http.Error(w, "Error during modification event", http.StatusInternalServerError)
		return
	}
}

func eventDelete(h *EventHandler, w http.ResponseWriter, uuid string) {
	if err := h.strg.Delete(uuid); err != nil {
		logger.Error(err)
		http.Error(w, "Error during deletion event", http.StatusInternalServerError)
		return
	}
}
