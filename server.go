package main

import (
	"fmt"
	"net/http"
	"strings"
)

type StationStore interface {
	GetStationURL(name string) string
}

type StationServer struct {
	store StationStore
}

func (s *StationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	stationName := strings.TrimPrefix(r.URL.Path, "/stations/")
	fmt.Fprint(w, s.store.GetStationURL(stationName))
}
