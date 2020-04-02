package main

import (
  "fmt"
)

func main() {
  var n,k int
  fmt.Scan(&n)
  fmt.Scan(&k)
  a := make([]int, n)
  for i:=0; i<n; i++ {
    fmt.Scan(&a[i])
  }
  var null [2]int
  fmt.Scan(&null[0])
  fmt.Scan(&null[1])

  // imos method
  imos := make([]int, n+k+1)
  for x:=0; x<k; x++ {
    imos[0] += a[x]
  }
  for j:=1; j<n-1; j++ {
    imos[j] = imos[j] + a[k] - a[j-1]
  }
  fmt.Println(imos)
}
