package main

import (
	// need import here for best looking output
	"time"
)

func main() {
	now := time.Now		// this doesn't work
	println("The current time is", now)
}
