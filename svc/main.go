package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jasontconnell/go-web-project/conf"
	"github.com/jasontconnell/go-web-project/handlers"
)

func main() {
	fn := flag.String("c", "config.json", "config file")
	flag.Parse()

	cfg, err := conf.LoadConfig(*fn)
	if err != nil {
		log.Fatal(err)
	}

	h := handlers.GetHandler(cfg)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("no port specified")
	}

	log.Fatal(http.ListenAndServe(":"+port, h))
}
