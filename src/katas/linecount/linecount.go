package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func CountLines( /* what should the param be? */ ) (int, error) {
	sc := bufio.NewScanner(r)
	var lines int
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

func main() {
	// How do you refer to stdin using the 'os' module?
	lines, err := CountLines( /* ... */ )
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
