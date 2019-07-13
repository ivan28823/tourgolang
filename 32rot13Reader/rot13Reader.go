package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {
	capital := x >= 'A' && x <= 'Z'
	if !capital && (x < 'a' || x > 'z') {
		return x // Not a letter
	}

	x += 13
	if capital && x > 'Z' || !capital && x > 'z' {
		x -= 26
	}
	return x
}
func (rd rot13Reader) Read(b []byte) (int, error) {
	n, err := rd.r.Read(b)

	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
