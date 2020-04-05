package main

import (
  "fmt"
  "sort"
)

func main() {
  var n,m,ans int
  fmt.Scanf("%d %d", &n, &m)

  a := make([]int, n)
  for i:=0; i<n; i++ {
    fmt.Scan(&a[i])
  }

  sort.Sort(sort.Reverse(sort.IntSlice(a)))

  m4 := float64(sum(a))/float64((4*m))
  for j:=0; j<m; j++ {
    x := a[:1][0]
    a = a[1:]
    if float64(x) >= m4 {
      ans++
    }
  }

  if ans == m {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
}

func sum(array []int) int {
  var buf int
  for _, v := range array {
    buf += v
  }
  return buf
}
