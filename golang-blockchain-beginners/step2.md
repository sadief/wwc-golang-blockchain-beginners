# Create A Genesis Block

The first thing we want to do when our server starts is create the initial block on the Blockchain. The `Genesis Block`.

We'll define the function to just print something out for now and call it by putting `createGenesisBlock()` inside of the `main` function.

Click the code below to replace the code in the editor so that we are calling `createGenesisBlock`.

<pre class="file" data-filename="main.go" data-target="replace">
package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	createGenesisBlock()

	log.Printf("Listening on port %v", port)
	http.ListenAndServe(port, nil)
}
</pre>

Now we'll stub out our `createGenesisBlock` function, showing a log for now.

<pre class="file" data-filename="main.go" data-target="append">
func createGenesisBlock() {
	log.Printf("Create first Block on the blockchain")
}
</pre>

Now when we run the server we should see `Create first Block on the blockchain` in the terminal.

`go run main.go`{{execute interrupt}}