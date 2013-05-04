package main

import (
  "fmt"
)

func main() {
  monthdays := map[string]int {
          "Jan": 31, "Feb": 28, "Mar": 31,
          "Apr": 30, "May": 31, "Jun": 30,
          "Jul": 31, "Aug": 31, "Sep": 30,
          "Oct": 31, "Nov": 30, "Dec": 31,
  }

  year := 0
  for _, days := range monthdays {
    year += days
  }

  fmt.Printf("Numbers of days in a year: %d\n", year)

  // Adding
  monthdays["Undecim"] = 30

  // Write over
  monthdays["Feb"] = 29

  // Check map for value

  v, ok := monthdays["Jan"] // ok is true if value exists
  if ok {
    fmt.Printf("Jan has value of %d\n", v)
  }

  // Removing
  delete(monthdays, "Mar")
}
