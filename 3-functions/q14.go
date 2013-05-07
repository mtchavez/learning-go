package main

import (
  "fmt"
)

func BubbleSort(list []int) []int {
  swapped := false
  for i := 1; i < len(list); i++ {
    if list[i-1] > list[i] {
      x, y := list[i-1], list[i]
      list[i-1] = y
      list[i] = x
      swapped = true
    }
  }

  if swapped {
    BubbleSort(list)
  }
  return list
}

func main() {
  list := []int {1,7,34,6,45,123,76,45,235,76,454,1,3,5,4,3,1}
  fmt.Printf("List before: %v\n", list)
  BubbleSort(list)
  fmt.Printf("List after: %v\n", list)
}
