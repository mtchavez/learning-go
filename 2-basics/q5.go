package main

import (
  "fmt"
)

func average(slice []float64) float64{
  size := float64(len(slice))
  if size == 0 {
    return float64(0.0)
  }
  var sum float64 = 0.0
  for _, val := range slice {
    sum += val
  }
  return sum / size
}

func main() {
  slice := []float64{1.9, 2.8, 3.7, 4.6, 5.5, 6.4, 7.3, 8.2, 9.1}
  fmt.Printf("Average of slice is: %f\n", average(slice))
}
