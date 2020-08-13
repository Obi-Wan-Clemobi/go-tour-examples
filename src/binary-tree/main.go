package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Exercise: Equivalent Binary Trees
// 1. Implement the Walk function.
// 2. Test the Walk function.
// The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.
// Create a new channel ch and kick off the walker:
// go Walk(tree.New(1), ch)
// Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.
// 3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.
// 4. Test the Same function.
// Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.
// The documentation for Tree can be found here.

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

// Need this so ch can close in Walk
func _walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		_walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		_walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for node1 := range ch1 {
		if node1 != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	ch := make(chan int, 100)
	go Walk(tree.New(1), ch)

	fmt.Println("Printing walk...")
	for node := range ch {
		fmt.Printf("%d ", node)
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Comparing tree 1 to tree 1")
	fmt.Println(Same(tree.New(1), tree.New(1)))

	fmt.Println()

	fmt.Println("Comparing tree 1 to tree 2")
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
