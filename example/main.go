package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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

type CountedVotes struct {
	Cats int
	Dogs int
}

var countedVotes CountedVotes

const port = ":3000"

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

	if strings.ToLower(v.Candidate) == "cats" {
		countedVotes.Cats++
	}

	if strings.ToLower(v.Candidate) == "dogs" {
		countedVotes.Dogs++
	}

	log.Printf("Votes are being calculated: \nCats: %v \nDogs: %v", countedVotes.Cats, countedVotes.Dogs)

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

// curl -X POST http://6c038964.ngrok.io -d '{"Voter": "Sadie", "Candidate": "Cats"}'
