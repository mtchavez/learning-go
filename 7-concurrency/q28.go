package main

import (
	"fmt"
)

func calc(in chan int) (chan int, chan int, chan int) {
	a, b, out := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			x := <-in
			a <- x
			b <- x
			out <- x
		}
	}()
	return a, b, out
}

func fib() chan int {
	x := make(chan int, 2)
	a, b, out := calc(x)
	go func() {
		x <- 0
		x <- 1
		for {
			_a := <-a
			_b := <-b
			x <- (_a + _b)
		}
	}()
	return out
}

func main() {
	x := fib()
	for i := 0; i < 10; i++ {
		fmt.Printf("output %d ", <-x)
	}
	fmt.Println()
}
