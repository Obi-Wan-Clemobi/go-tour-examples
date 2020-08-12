package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	reader io.Reader
}

func (rot rot13Reader) Read(bytes []byte) (n int, err error) {
	n, err = rot.reader.Read(bytes)
	for i := range bytes {
		if (bytes[i] >= 'A' && bytes[i] < 'N') || (bytes[i] >= 'a' && bytes[i] < 'n') {
			bytes[i] += 13
		} else if (bytes[i] > 'M' && bytes[i] <= 'Z') || (bytes[i] > 'm' && bytes[i] <= 'z') {
			bytes[i] -= 13
		}
	}
	return
}

func main() {
	msg := "Lbh penpxrq gur pbqr!"
	s := strings.NewReader(msg)
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
