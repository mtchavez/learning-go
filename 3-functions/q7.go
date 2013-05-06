package main

import (
  "fmt"
  "sort"
)

type IntSlice []int

// Sort interface functions

func (x IntSlice) Len() int{
  return len(x)
}

func (x IntSlice) Less(i, j int) bool {
  return x[i] < x[j]
}

func (x IntSlice) Swap(i, j int) {
  x[i], x[j] = x[j], x[i]
}

func orderedParams(args ...int) {
  sort.Sort(IntSlice(args))
  fmt.Println(args)
}

func main() {
  orderedParams(8,3)
  orderedParams(3,8)
}
