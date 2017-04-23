package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Make sure there's a URL on the command line
	// Retrieve the content from the URL in Args[1]
	if err != nil {
		log.Fatal(err)
	}
	// Close the http.Response object
	io.Copy(os.Stdout, resp.Body)
}
