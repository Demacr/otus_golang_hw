package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/logger"
)

func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func writeJSON(w http.ResponseWriter, value interface{}) {
	data, err := json.Marshal(&value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error("failed to marshal:", err)
		fmt.Fprintf(w, "failed to marshal: %v", err)
		return
	}

	if _, err = w.Write(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
