package main

import (
	"log"
	"net/http"

)"github.com/davecgh/go-spew/spew"

const port = ":8080"

func main() {
	log.Printf("Listening on port %v", port)
	http.ListenAndServe(port, nil)
}
