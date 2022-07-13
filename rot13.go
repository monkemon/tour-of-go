// not really proud of this one since i copied from wikipedia but it gets the job done and i spent little time on it so i call it being efficient

// solution to https://go.dev/tour/methods/23

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// copy paste from wiki :)
func rot (b byte) byte {
	if b >= 'A' && b <= 'Z' {
			b = 'A' + (b - 'A' + 13) % 26
		} else if b >= 'a' && b <= 'z' {
			b = 'a' + (b - 'a' + 13) % 26
		}
	return b
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
    n, err = r.r.Read(p)
    for i := 0; i < n; i++ {
        p[i] = rot(p[i])
    }
    return n, err
}


func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
