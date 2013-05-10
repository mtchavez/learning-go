package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type T struct {
	a int
}

func Set(t *T) {
	fmt.Println(t)
}

func Set2(t T) {
	fmt.Println(&t)
}

func main() {
	var p1 Person
	p2 := new(Person)
	fmt.Println(p1)
	fmt.Println(p2)

	t := T{a: 5}
	Set(&t)
	Set2(t)
}
