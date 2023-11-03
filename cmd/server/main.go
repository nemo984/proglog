package main

import (
	"log"

	"github.com/nemo984/dislog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
