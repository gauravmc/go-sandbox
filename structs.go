package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{5, 6}
  fmt.Println(Vertex{5, 6})
  fmt.Println(v.X)
}