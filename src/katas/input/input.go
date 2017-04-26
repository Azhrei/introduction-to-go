package fred

import (
	"bufio"
	"io"
	"log"
	"strings"
)

// ReadAll reads all the lines of text from r and returns
// all the data read as a string
func ReadAll(r io.Reader) string {
	sc := bufio.NewScanner(r)
	var lines []string // = make([]string, 0, 100)
	for sc.Scan() {
		// Is appending one item at a time a good idea?
		lines = append(lines, sc.Text())
	}
	err := sc.Err()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Join(lines, "\n")
}
