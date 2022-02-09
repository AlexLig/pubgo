package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETStations(t *testing.T) {
	t.Run("returns off radio's url", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/stations/OffRadio", nil)
		response := httptest.NewRecorder()

		StationServer(response, request)

		got := response.Body.String()
		want := "https://s9.yesstreaming.net:17008/stream"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
