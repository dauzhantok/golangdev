package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	switch {
	case x == 0 || x == 1:
		return x
	case x < 0:
		return 0
	}
	z := x
	t := 0
	for {
		z -= (z*z - x) / (2 * z)
		if z*z == x || t == 100 {
			break
		}
		t++
	}

	return z
}

func main() {
	fmt.Println(Sqrt(4))
}
