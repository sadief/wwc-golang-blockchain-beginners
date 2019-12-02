# Refactor createGenesisBlock function

We're going to slightly refactor our `createGenesisBlock` to use the `createHash` function.

```go
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
```

Clcik to add the code to your editor.

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
</pre>

`go run main.go`{{execute interrupt T1}}

Now check out the hash when we make a GET request.

`curl localhost:8080`{{execute interrupt T2}}