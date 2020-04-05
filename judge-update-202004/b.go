package main

import (
  "fmt"
  "sort"
)

func main() {
  var (
    n,x int
    c string
    red,blue []int
  )
  fmt.Scan(&n)
  for i:=0; i<n; i++ {
    fmt.Scan(&x)
    fmt.Scan(&c)
    switch(c) {
    case "R":
      red = append(red, x)
    case "B":
      blue = append(blue, x)
    }
  }
  sort.Sort(sort.IntSlice(red))
  sort.Sort(sort.IntSlice(blue))
  for _, r := range red {
    fmt.Println(r)
  }
  for _, b := range blue {
    fmt.Println(b)
  }
}
