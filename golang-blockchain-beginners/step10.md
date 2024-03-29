# Create Vote struct

We want to parse our request message into a `Vote` struct, so first we define it:

```go
type Vote struct {
	Voter     string
	Candidate string
}
```

Click to copy the code into your editor.

<pre class="file" data-filename="main.go" data-target="replace">
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Vote struct {
	Voter     string
	Candidate string
}

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getBlockchain(w, r)
		case "POST":
			addNewBlock(w, r)
		}
	})

	log.Printf("Listening on port %v", port)
	http.ListenAndServe(port, nil)
}

func createGenesisBlock() {
	t := time.Now()
	genesisBlock := Block{
		Timestamp: t.String(),
		Voter:     "",
		Candidate: "",
	}
	genesisBlock.Hash = createHash(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)
	log.Printf("Create first Block on the blockchain")
}

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	log.Printf("Blockchain: %#v", Blockchain)
}

func addNewBlock(w http.ResponseWriter, r *http.Request) {
	log.Printf("Add new blockchain to the block")
}

func createHash(b Block) string {
	record := fmt.Sprintf("%v%v%v%v", b.Timestamp, b.Voter, b.Candidate, b.PrevHash)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func createNewBlock(prevBlock Block, Voter, Candidate string) Block {
	t := time.Now()
	newBlock := Block{
		Timestamp: t.String(),
		Voter:     Voter,
		Candidate: Candidate,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.Hash = createHash(newBlock)
	return newBlock
}
</pre>

Let's continue on to step 11.