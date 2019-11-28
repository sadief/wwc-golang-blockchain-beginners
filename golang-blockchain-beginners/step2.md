# Create A Genesis Block

<pre class="file" data-filename="main.go" data-target="replace">
package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	createGenesisBlock()

	log.Printf("Listening on port %v", port)
	http.ListenAndServe(port, nil)
}

func createGenesisBlock() {
	log.Printf("Create first Block on the blockchain")
}
</pre>

`go run main.go`{{execute}}