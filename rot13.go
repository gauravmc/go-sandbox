package main

import (
  "io"
  "os"
  "strings"
)

func rotate_by_13_places(b byte) byte {
  // pending
  return b
}

type rot13Reader struct {
  r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
  n, err = r.r.Read(p)
  for i := 0; i < n; i++ {
    p[i] = rotate_by_13_places(p[i])
  }
  return
}

func main() {
  s := strings.NewReader("abcDEFghijKLmnopQRStuVWXYz")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}
