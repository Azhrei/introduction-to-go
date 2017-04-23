package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := // Retrieve the data...
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	result := // Make the map variable
	dec := // Create decoder/parser object
	err := // Convert the JSON into a Go object
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("My IP address is:", result["origin"])
}
