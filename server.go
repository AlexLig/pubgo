package main

import (
	"fmt"
	"net/http"
	"strings"
)

func StationServer(w http.ResponseWriter, r *http.Request) {
	stationName := strings.TrimPrefix(r.URL.Path, "/stations/")
	fmt.Fprint(w, getStationURL(stationName))
}

func getStationURL(name string) string {
	if name == "EnLefko" {
		return "EnLefkoURL"
	}
	if name == "OffRadio" {
		return "OffRadioURL"
	}

	return ""
}
