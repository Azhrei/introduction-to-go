package count

import (
	"bufio"
	"log"
	"os"
)

// CountLines opens the file referenced by the parameter, then
// counts the number of lines therein. Last, it returns the count.
func CountLines(path string) int {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(f)
	var lines int
	// Read one line at a time in a loop, incrementing 'lines'
	// ...
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	f.Close()
	return lines
}
