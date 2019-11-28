# Put the Block in the Chain

<pre class="file" data-filename="main.go" data-target="replace">
package main

import (
	"log"
	"net/http"
)

type Block struct {
	Hash      string
	PrevHash  string
	Timestamp string
	Voter     string
	Candidate string
}

var Blockchain []Block

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