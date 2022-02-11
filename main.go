package main

import (
	"log"
	"net/http"
	"pubgo/server"
)

func main() {
	server := server.Init()
	log.Fatal(http.ListenAndServe(":8080", server))
}
