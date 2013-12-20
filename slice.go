package main

import "fmt"

func main() {
  slice := []int{1, 2, 3, 4, 5}
  fmt.Println(slice)

  for i := 0; i < 4; i++ {
    fmt.Printf("slice[%d] == %d\n", i, slice[i])
  }

  fmt.Println("slice[1:3] ==", slice[1:3])
}
