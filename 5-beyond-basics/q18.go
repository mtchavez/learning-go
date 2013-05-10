package main

import (
	"fmt"
)

func main() {
	p := 0
	fmt.Printf("p is %d\n", p)
	q := &p
	*q = p + 1
	fmt.Printf("p is %d\n", p)
}
