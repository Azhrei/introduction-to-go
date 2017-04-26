package readers

import (
	"io"
)

// Combine returns an io.Reader which represents
// the contents of a and b.
func Combine(a, b io.Reader) io.Reader {
	// Which piece of the io package returns a Reader that starts
	// reading from a second Reader when the first one is exhausted?
	return io.MultiReader(a, b)
}

// always reader always fills the read buffer with
// the byte ch.
type alwaysReader struct {
	ch byte
}

func (a *alwaysReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = a.ch
	}
	return len(buf), nil
}

// AReader returns an io.Reader which returns n 'A' characters
func AReader(n int) io.Reader {
	r := &alwaysReader{'A'}
	return io.LimitReader(r, int64(n))
}
