package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// select is acting on which channel has activity and then eval the right line
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// make a channel of int
	c := make(chan int)
	// make another channel of int
	quit := make(chan int)
	go func() {
		// print values from c as they come
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		// after we've received all 10 fibs, send 0 to quit channel
		quit <- 0
	}()
	fibonacci(c, quit)
}
