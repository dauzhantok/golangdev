
import (
	"fmt"
)

type ErrNegativeSqrt float64 

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	switch {
	case x == 0 || x == 1:
		return x, nil
	case x < 0:
		return x, ErrNegativeSqrt(x)
		
	}
	z := x
	t := 0
	for {
		z -= (z*z - x) / (2*z)
		if z*z==x || t==100{
			break
		}
		t++
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}