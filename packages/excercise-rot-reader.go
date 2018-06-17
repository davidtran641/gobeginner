package main
import (
 	"io"
 	"os"
 	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	in := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz "
	out := "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm "
	charMap := make(map[byte]byte)
	for i := 0; i < len(in); i++ {
		charMap[in[i]] = out[i]
	}

	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = charMap[b[i]]
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}