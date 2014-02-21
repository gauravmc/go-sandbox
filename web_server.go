package main

import "fmt"
import "net/http"

type Hello{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World!")
}

func main() {
  var h Hello
  http.ListenAndServe("localhost:4000", h)
}
