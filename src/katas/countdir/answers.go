// +build ignore

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// CountLines reads from the io.Reader parameter, returning an integer
// describing the number of lines read in.
func CountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	var lines int
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

// CountFile opens the file referenced by the parameter, then calls
// CountLines() and prints its result.
func CountFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	lines, err := CountLines(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\t%d\n", path, lines)
}

// CountDir opens the directory name provided as a parameter, then calls
// CountFile() for each directory entry that is an ordinary file whose
// name ends with '.txt'.
func CountDir(dir string) {
	d, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()
	entries, err := d.Readdir(-1)
	if err != nil {
		// Error could be 'not a directory'
		// or 'no permission' or ...?
		log.Fatal(err)
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		if strings.HasSuffix(e.Name(), ".txt") {
			CountFile(filepath.Join(dir, e.Name()))
		}
	}
}

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		CountDir(arg)
	}
}
