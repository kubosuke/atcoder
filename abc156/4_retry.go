package main

import (
  "fmt"
)

var mod = int64()

func main() {
  var (
    n,a,b int64
    result int64
  )
  fmt.Scan(&n)
  fmt.Scan(&a)
  fmt.Scan(&b)

  result = power(2,n) - 1 - nCk(n,a) - nCk(n,b)
  for result < 0 {
    result += mod
  }

  fmt.Println(result)
}

func power(a,b int64) int64 {
  var ret int64 = 1
  for b > 0 {

  }
}
