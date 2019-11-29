# Put the Block in the Chain

Let's add a `Block` struct which contains the fields we want on our Blockchain, and then define a `Blockchain` which is a slice of Blocks.

We'll add the following to our `main.go` right after our imports

```
type Block struct {
	Hash      string
	PrevHash  string
	Timestamp string
	Voter     string
	Candidate string
}

var Blockchain []Block
```

Click the code to change the code in our editor to contain `Block` and `Blockchain`.

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

We won't see any changes by running the server just yet, so let's continue to step 4.