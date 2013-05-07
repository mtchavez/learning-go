package even

import (
  "testing"
  "fmt"
)

func TestEven(t *testing.T) {
  if ! Even(2) {
    t.Errorf("2 should be even")
  }

}

func ExampleEven() {
  if Even(2) {
    fmt.Printf("Is even\n")
  }
  // Output:
  // Is even
}
