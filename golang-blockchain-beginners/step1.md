# Spin up a Go Server

```go
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
```

`echo "go run main.go"`{{execute}}