// +build ignore

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// CountLines reads from the io.Reader parameter and returns the number
// of lines contained therein.
func CountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	var lines int
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

func main() {
	lines, err := CountLines(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
