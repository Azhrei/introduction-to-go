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

func CountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	var lines int
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

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
		// Skip if 'e' is a directory
		// Skip unless 'e' ends with '.txt'
		// Count and print the number of lines
		//   (be careful with the pathname!)
	}
}

func main() {
	args := // How do you get data from the command line?
	for _, arg := range args {
		CountDir(arg)
	}
}
