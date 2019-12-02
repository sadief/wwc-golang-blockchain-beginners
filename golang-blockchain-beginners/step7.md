# Write createHash function


Now we'll implement a function to create the has for each of our blocks. This will take in a Block, concatenate the data, and hash it using the Go "crypto" library which we import in with `"crypto/sha256"`

```go
func createHash(b Block) string {
	record := fmt.Sprintf("%v%v%v%v", b.Timestamp, b.Voter, b.Candidate, b.PrevHash)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
```

Click the code below to copy to your editor.

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
		Hash:      "",
		Timestamp: t.String(),
		Voter:     "",
		Candidate: "",
	}
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

We aren't doing anything with this just yet, so let's head to step 8.