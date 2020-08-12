package main

import (
	"fmt"
	"math"
)

const threshold = 1e-6

func Sqrt(x float64) float64 {
	z := x
	n := 0.0
	for math.Abs(n-z) > threshold {
		n, z = z, z-(z*z-x)/(2*z)
	}
	return z
}

func main() {
	fmt.Println("Flow control exercise")
	fmt.Printf("Math sqrt: %f", math.Sqrt(2))
	fmt.Println()
	fmt.Printf("My sqrt  : %f", Sqrt(2))
	fmt.Println()
}
