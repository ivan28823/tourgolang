package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	size := len(b)
	for t := 0; t < size; t++ {
		b[t] = 'A'
	}
	return size, nil
}

func main() {
	reader.Validate(MyReader{})
}
