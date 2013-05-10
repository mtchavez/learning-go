package main

import (
	"fmt"
)

func list_map(list []int, mapFun func(int) int) []int {
	for i, item := range list {
		list[i] = mapFun(item)
	}
	return list
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	f := func(i int) int {
		return i * 3
	}
	fmt.Printf("List Before: %v\n", list)
	fmt.Printf("List After Map: %v\n", list_map(list, f))
}
