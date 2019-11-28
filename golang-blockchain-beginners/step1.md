# Spin up a Go Server

`touch main.go`{{execute}}

<pre class="file" data-filename="main.go" data-target="append">
package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	log.Printf("Listening on port %v", port)
	http.ListenAndServe(port, nil)
}
</pre>

`go run main.go`{{execute}}