package main

import "fmt"
import "math"

func Sqrt(x float64) float64 {
  z := 1.0
  var n float64
  for math.Abs(z - n) > 0.001 {
    n = z
    z = z - ((z*z - x)/2*z)
  }
  return z
}

func main() {
  fmt.Println(Sqrt(2))
}
