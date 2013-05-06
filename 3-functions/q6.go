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
  slice := []float64{1.1, 4.5, 82.3, 4.0, 62.3}
  fmt.Printf("Average: %f\n", average(slice))
}
