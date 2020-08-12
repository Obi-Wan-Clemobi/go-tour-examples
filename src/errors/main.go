package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("Complex numbers not supported, input=%f", float64(e))
}

const threshold = 1e-6

func Sqrt(x float64) (float64, error) {
	if x > 0.0 {
		z := x
		n := 0.0
		for math.Abs(n-z) > threshold {
			n, z = z, z-(z*z-x)/(2*z)
		}
		return z, nil
	}

	return 0, ErrNegativeSqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
