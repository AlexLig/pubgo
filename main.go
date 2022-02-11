package main

import (
	"log"
	"net/http"
)

func main() {
	server := &StationServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
