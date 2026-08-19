package main

import (
	"log"

	httpserver "DCacheX/internal/server/http"
)

func main() {
	server := httpserver.NewServer(":8080")
	log.Println("Http server listening on :8080")

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
