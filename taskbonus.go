
import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (ro *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = ro.r.Read(b)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'M' || b[i] >= 'a' && b[i] <= 'm' {
			b[i] = b[i] + 13
		} else {
			b[i] = b[i] - 13
		}
	}
	return n, err

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}