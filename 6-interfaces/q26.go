package main

import (
	"fmt"
)

func Greater(x, y interface{}) bool {
	switch x.(type) {
	case int:
		if _, ok := y.(int); ok {
			return x.(int) > y.(int)
		}
	case float32:
		if _, ok := y.(float32); ok {
			return x.(float32) > y.(float32)
		}
	}
	return false
}

func main() {
	fmt.Printf("Is %d, greater than %d: %v\n", 5, 4, Greater(5, 4))
	fmt.Printf("Is %d, greater than %d: %v\n", 4, 4, Greater(4, 4))
	fmt.Printf("Is %d, greater than %d: %v\n", 3, 4, Greater(3, 4))
}
