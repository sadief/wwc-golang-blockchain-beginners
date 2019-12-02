# Implement getBlockchain Function

Now we are going to use a pretty-print library, `Spew` to print out our Blockchain when a GET request is made. We'll update the `getBlockchain` function to look like this

```go
func getBlockchain(w http.ResponseWriter, r *http.Request) {
	log.Printf("Blockchain: %#v", Blockchain)
}
```

Click the code below to copy it over to the editor with the updated function.

<pre class="file" data-filename="main.go" data-target="replace">
package main

import (
	"log"
	"net/http"
	"time"
	"github.com/davecgh/go-spew/spew"
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
		PrevHash:  "",
		Timestamp: t.String(),
		Voter:     "",
		Candidate: "",
	}
	Blockchain = append(Blockchain, genesisBlock)
	log.Printf("Create first Block on the blockchain")
}

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	spew.Dump("Blockchain: %#v", Blockchain)
}

func addNewBlock(w http.ResponseWriter, r *http.Request) {
	log.Printf("Add new blockchain to the block")
}
</pre>

Now we should be able to make a get request and see our Blockchain with it's first block printed out.

`go run main.go`{{execute interrupt}}
`curl localhost:8080`{{execute interrupt}}