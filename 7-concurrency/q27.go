package main

import (
	"fmt"
)

func PrintMe(ch chan int, quit chan bool) {
	for {
		select {
		case x := <-ch:
			fmt.Println(x)
		case <-quit:
			break
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go PrintMe(ch, quit)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	quit <- true
}
