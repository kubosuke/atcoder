package main

import (
  "fmt"
)

func main() {
  var (
    n,k int
    d int
    tmp int
    res int
  )
  fmt.Scan(&n)
  fmt.Scan(&k)
  ans := make([]int, n)

  for i:=0; i<k; i++ {
    fmt.Scan(&d)
    for j:=0; j<d; j++ {
      fmt.Scan(&tmp)
      ans[tmp-1] = 1
    }
  }
  for _, x := range ans {
    res += x
  }
  fmt.Println(n-res)
}
