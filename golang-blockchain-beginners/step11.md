# Implement addNewBlock function

Lastly, we'll implement the `addNewBlock` that we currently have stubbed out. This will handle the incoming POST request, and decode the body into the Vote struct. Then it will grab the last block, create a new one, and append it to the Blockchain.

```go
func addNewBlock(w http.ResponseWriter, r *http.Request) {
	var v Vote
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&v)

	if err != nil {
		log.Printf("Error Decoding Body")
	}
	defer r.Body.Close()

	lastBlock := Blockchain[len(Blockchain)-1]
	newBlock := createNewBlock(lastBlock, v.Voter, v.Candidate)
	Blockchain = append(Blockchain, newBlock)
}
```

Click to add this code to the editor.

<pre class="file" data-filename="main.go" data-target="replace">
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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
	var v Vote
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&v)

	if err != nil {
		log.Printf("Error Decoding Body")
	}
	defer r.Body.Close()

	lastBlock := Blockchain[len(Blockchain)-1]
	newBlock := createNewBlock(lastBlock, v.Voter, v.Candidate)
	Blockchain = append(Blockchain, newBlock)
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


Congratulations! You've made a Blockchain! Let's test it out:

`go run main.go`{{execute interrupt T1}}

Remember to see the logs you'll need to flip back into Terminal 1.
`curl localhost:8080`{{execute T2}}

Pop in your name and vote for cats or dogs!
`curl -X POST localhost:8080 -d '{"Voter": "Name", "Candidate": "CatsorDogs"}'`{{execute T2}}

And now let's have a look at the new block we created.
`curl localhost:8080`{{execute T2}}