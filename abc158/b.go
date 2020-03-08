package main

import (
  "fmt"
)

func main() {
  var (
    a,b int
  )
  fmt.Scan(&a)
  fmt.Scan(&b)

  fmt.Println(check(a,b))
}

func check(a,b int) int {
  for i:=8; i<10000; i++ {
    if (i*8)/100 == a && i/10 == b{
      return i
    }
  }
  return -1
}
