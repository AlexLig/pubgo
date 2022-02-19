package main

import (
	"github.com/AlexLig/pubgo/server"
	"log"
	"net/http"
)

func main() {
	server := server.Init()
	log.Fatal(http.ListenAndServe(":8080", server))
}
