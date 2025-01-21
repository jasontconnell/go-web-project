package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jasontconnell/go-web-project/handlers"
)

func main() {
	devmode := os.Getenv("devmode") == "true"
	h := handlers.GetHandler(devmode)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("no port specified")
	}

	log.Fatal(http.ListenAndServe(":"+port, h))
}
