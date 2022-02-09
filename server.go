package main

import (
	"fmt"
	"net/http"
)

func StationServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "https://s9.yesstreaming.net:17008/stream")
}
