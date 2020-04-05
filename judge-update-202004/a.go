package main

import (
  "fmt"
)

func main() {
  var s,l,r int
  fmt.Scan(&s)
  fmt.Scan(&l)
  fmt.Scan(&r)

  if s<l {
    fmt.Println(l)
  } else if r<s {
    fmt.Println(r)
  } else {
    fmt.Println(s)
  }
}
