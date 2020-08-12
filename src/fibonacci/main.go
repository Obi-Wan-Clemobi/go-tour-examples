package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f0, f1, sum := 0, 0, 0
	return func() int {
		if f1 == 0 {
			f1 = 1
		} else {
			sum, f0 = f0+f1, f1
			f1 = sum
		}
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", f())
	}
}
