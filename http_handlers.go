package main

import "fmt"
import "net/http"

type String string

type Struct struct{
  Greeting string
  Punct string
  Who string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%s", s)
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%s%s %s", s.Greeting, s.Punct, s.Who)
}

func main() {
  http.Handle("/string", String("I'm a String"))
  http.Handle("/struct", Struct{"Hi", ",", "Struct"})
  http.ListenAndServe("localhost:4000", nil)
}
