package httpserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
)

var tt_invalid_pathes = []struct {
	Path string
}{
	{
		"/health",
	},
	{
		"/123",
	},
	{
		"/api",
	},
}

var tt_invalid_pathes_api = []struct {
	Path string
}{
	{
		"/api/",
	},
	{
		"/api/test",
	},
}

var tt_api_event_add = []struct {
	Request *AddRequest
}{
	{
		&AddRequest{
			"UUID-1",
			"Some event 1",
			time.Date(2020, 11, 11, 12, 0, 0, 0, time.Local),
			time.Hour,
			"very important event",
			"User-1",
			time.Hour + 24,
		},
	},
}

var tt_api_event_modify = []struct {
	Request *ModifyRequest
	UUID    string
}{
	{
		UUID: "UUID-1",
		Request: &ModifyRequest{
			"Some event 1",
			time.Date(2020, 11, 11, 13, 0, 0, 0, time.Local), // Move start to 13:00
			time.Hour,
			"very important event",
			"User-1",
			time.Hour + 24,
		},
	},
}

var tt_api_event_delete = []struct {
	UUID string
}{
	{
		"UUID-1",
	},
}

var tt_api_event_list = []struct {
	Type     string
	Datetime string
	Expected int
}{
	{
		"day",
		"2020-11-11",
		1,
	},
	{
		"week",
		"2020-11-09",
		1,
	},
	{
		"month",
		"2020-11-01",
		1,
	},
}

func TestHTTPServerSimple(t *testing.T) {
	strg := storage.NewInMemoryStorage()
	router := NewRouter(strg)

	srv := httptest.NewServer(router.RootHandler())
	defer srv.Close()

	t.Run("check root path", func(t *testing.T) {
		res, err := http.Get(srv.URL + "/")
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("check invalid pathes", func(t *testing.T) {
		for _, tt := range tt_invalid_pathes {
			res, err := http.Get(srv.URL + tt.Path)
			require.NoError(t, err)

			require.Equal(t, http.StatusNotFound, res.StatusCode, "path = "+tt.Path)
		}
	})

	t.Run("check /api pathes", func(t *testing.T) {
		for _, tt := range tt_invalid_pathes_api {
			res, err := http.Get(srv.URL + tt.Path)
			require.NoError(t, err)

			require.Equal(t, http.StatusNotFound, res.StatusCode, "path = "+tt.Path)
		}
	})

	t.Run("check /api/event POST", func(t *testing.T) {
		for _, tt := range tt_api_event_add {
			payload := bytes.Buffer{}
			err := json.NewEncoder(&payload).Encode(tt.Request)
			require.NoError(t, err)

			res, err := http.Post(srv.URL+"/api/event", "application/json", &payload)
			require.Equal(t, http.StatusOK, res.StatusCode)
		}
	})

	t.Run("check /api/event PATCH", func(t *testing.T) {
		for _, tt := range tt_api_event_modify {
			payload := bytes.Buffer{}
			err := json.NewEncoder(&payload).Encode(tt.Request)
			require.NoError(t, err)

			client := http.Client{}
			req, err := http.NewRequest("PATCH", srv.URL+"/api/event/"+tt.UUID, &payload)
			require.NoError(t, err)
			res, err := client.Do(req)
			require.NoError(t, err)

			require.Equal(t, http.StatusOK, res.StatusCode)
		}
	})

	t.Run("check /api/event DELETE", func(t *testing.T) {
		for _, tt := range tt_api_event_delete {
			client := http.Client{}
			req, err := http.NewRequest("DELETE", srv.URL+"/api/event/"+tt.UUID, nil)
			require.NoError(t, err)
			res, err := client.Do(req)
			require.NoError(t, err)

			require.Equal(t, http.StatusOK, res.StatusCode)
		}
	})
}

func TestHTTPServerComplex(t *testing.T) {
	strg := storage.NewInMemoryStorage()
	router := NewRouter(strg)

	srv := httptest.NewServer(router.RootHandler())
	defer srv.Close()

	for _, tt := range tt_api_event_add {
		payload := bytes.Buffer{}
		err := json.NewEncoder(&payload).Encode(tt.Request)
		require.NoError(t, err)

		res, err := http.Post(srv.URL+"/api/event", "application/json", &payload)
		require.Equal(t, http.StatusOK, res.StatusCode)
	}

	t.Run("Test List APIs", func(t *testing.T) {
		for _, tt := range tt_api_event_list {
			res, err := http.Get(fmt.Sprintf("%v/api/event?type=%v&dt=%v", srv.URL, tt.Type, tt.Datetime))
			require.NoError(t, err)
			events := []storage.Event{}

			err = json.NewDecoder(res.Body).Decode(&events)
			require.NoError(t, err)

			require.Equal(t, tt.Expected, len(events))
		}
	})
}
