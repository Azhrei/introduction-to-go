package errorhandling

import (
	"bufio"
	"os"
)

// Note the new return value!
func CountLines(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		// return what?
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var lines int
	for sc.Scan() {
		lines++
	}
	if err := sc.Err(); err != nil {
		// return what?
	}
	// return what?
}
