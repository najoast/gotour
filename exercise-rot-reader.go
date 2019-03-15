package main

import (
	"fmt"
	"io"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func isAlphabetical(c byte) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
		return true
	}
	return false
}

func rot13Decrypt(c byte) byte {
	if c >= 'A' && c <= 'M' {
		return c + 13
	} else if c >= 'N' && c <= 'Z' {
		return c - 13
	} else if c >= 'a' && c <= 'm' {
		return c + 13
	} else if c >= 'n' && c <= 'z' {
		return c - 13
	} else {
		return c
	}
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	fmt.Println("rot13Reader.Read", b)
	size, err := rot13.r.Read(b)
	fmt.Println("rot13Reader.Read", b)

	for i := 0; i < size; i++ {
		if isAlphabetical(b[i]) {
			b[i] = rot13Decrypt(b[i])
		}
	}
	return size, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	b := make([]byte, 32)
	r.Read(b)
	fmt.Printf("%q", b)
}
