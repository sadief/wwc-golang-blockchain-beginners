# Implement createGenesisBlock Function

Instead of logging some text, we want our `createGenesisBlock` function to actually do something. Here we're setting the timestamp to now, creating a new instance of the `Block` struct, and appending it to our `Blockchain`.

```go
func createGenesisBlock() {
	t := time.Now()
	genesisBlock := Block{
		Hash:      "",
		PrevHash:  "",
		Timestamp: t.String(),
		Voter:     "",
		Candidate: "",
	}
	Blockchain = append(Blockchain, genesisBlock)
	log.Printf("Blockchain: %#v", Blockchain)
	log.Printf("Create first Block on the blockchain")
```

Click on the code below to replace it with our filled in `createGenesisBlock` function

<pre class="file" data-filename="main.go" data-target="replace">
package main

import (
	"log"
	"net/http"
	"time"
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
	t := time.Now()
	genesisBlock := Block{
		Hash:      "",
		PrevHash:  "",
		Timestamp: t.String(),
		Voter:     "",
		Candidate: "",
	}
	Blockchain = append(Blockchain, genesisBlock)
	log.Printf("Blockchain: %#v", Blockchain)
	log.Printf("Create first Block on the blockchain")
}
</pre>

`go run main.go`{{execute interrupt}}