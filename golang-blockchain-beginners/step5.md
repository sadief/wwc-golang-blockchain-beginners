# Handle GET and POST requests

Now we want to handle two requests - GET and POST. Get will print out the Blockchain as it is currently, and POST will add a new Block to the Blockchain, using the data in the request body. We're going to change the server code to look like this:

```go
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
```

Click to replace the code in your editor with this code that includes the request handlers.

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
	log.Printf("Blockchain: %#v", Blockchain)
	log.Printf("Create first Block on the blockchain")
}
</pre>

And now let's stub out these two functions

<pre class="file" data-filename="main.go" data-target="append">
func getBlockchain(w http.ResponseWriter, r *http.Request) {
	log.Printf("Print out the Blockchain")
}

func addNewBlock(w http.ResponseWriter, r *http.Request) {
	log.Printf("Add new blockchain to the block")
}
</pre>


Now we can test each of these functions by starting the server and pinging it. If you flip back to `Terminal 1` you should see the log statements for each curl request.

`go run main.go`{{execute interrupt T1}}

`curl localhost:8080`{{execute T2}}

`curl -X POST localhost:8080`{{execute T2}}