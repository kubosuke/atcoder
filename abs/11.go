package main

import (
	"fmt"
	"strings"
)

func main() {
  var n int
	fmt.Scan(&n)
  t := make([]int, n)
  x := make([]int, n)
  y := make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan("%d %d %d", &t[i], &x[i], &y[i])
  }

}
