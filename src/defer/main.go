package main

import (
	"fmt"
)

func rule1() {
	i := 0
	// 1. A deferred function's arguments are evaluated when the defer statement is evaluated.
	defer fmt.Println("----> ", i)
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	return
}

func rule2() {
	i := 0
	// 2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.

	for ; i < 10; i++ {
		defer fmt.Println(i)
	}

	return
}

func rule3() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	fmt.Println("Rule 1: A deferred function's arguments are evaluated when the defer statement is evaluated.")
	rule1()
	fmt.Println("Rule 2: Deferred function calls are executed in Last In First Out order after the surrounding function returns.")
	rule2()
	fmt.Println("Rule 3: Deferred functions may read and assign to the returning function's named return values.")
	i := rule3()
	fmt.Println(i)
}
