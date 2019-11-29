# Spin up a Go Server

You can click on the commands in grey to run them, and you can click on the code to automatically copy it over into your terminal.

First thing we want to do is create a `main.go` file that we will run our code through.

`touch main.go`{{execute}}

When the `main.go` file appears in your editor, click into it, and then you can click on the following code snippet to copy it in (or feel free to type it out).

<pre class="file" data-filename="main.go" data-target="append">
package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	log.Printf("Listening on port %v", port)
	http.ListenAndServe(port, nil)
}
</pre>

Let's see this thing work - click on the command below and watch your terminal logs. You should see
`Listening on port :8080`

`go run main.go`{{execute}}