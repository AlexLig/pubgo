package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETStations(t *testing.T) {
	t.Run("returns off radio's url", func(t *testing.T) {
		request := newGetStationRequest("OffRadio")
		response := httptest.NewRecorder()

		StationServer(response, request)

		assertResponseBody(t, response.Body.String(), "OffRadioURL")

	})

	t.Run("returns en lefko url", func(t *testing.T) {
		request := newGetStationRequest("EnLefko")
		response := httptest.NewRecorder()

		StationServer(response, request)

		assertResponseBody(t, response.Body.String(), "EnLefkoURL")
	})
}

func newGetStationRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/stations/%s", name), nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
