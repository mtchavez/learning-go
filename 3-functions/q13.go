package main

import (
  "fmt"
)

func max(list []int) (m int) {
  m = list[0]
  for _, item := range list {
    if item > m {
      m = item
    }
  }
  return
}

func min(list []int) (m int) {
  m = list[0]
  for _, item := range list {
    if item < m {
      m = item
    }
  }
  return
}

func main() {
  list := []int {2891, 12, 34343, 843, 3843, 4398493}
  fmt.Printf("Max in list %v is %d\n", list, max(list))
  fmt.Printf("Min in list %v is %d\n", list, min(list))
}
