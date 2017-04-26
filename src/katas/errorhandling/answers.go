// +build ignore

package errorhandling

import (
	"bufio"
	"os"
)

// CountLines opens the file referenced by the parameter, then
// counts the number of lines therein. Last, it returns the count.
func CountLines(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var lines int
	for sc.Scan() {
		lines++
	}
	if err := sc.Err(); err != nil {
		return 0, err
	}
	return lines, nil
}
