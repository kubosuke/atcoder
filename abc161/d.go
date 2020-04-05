package main

import (
  "fmt"
)

func main() {
  var k int
  fmt.Scan(&k)
  fmt.Println(task(k))
}

func task(k int) int {
  var p int
  slice := []int{1,2,3,4,5,6,7,8,9}
  for i:=0; i < k; i++ {
    p = slice[i]
    if p%10 != 0 {
      slice = append(slice, 10*p + p%10 -1)
    }
    slice = append(slice, 10*p + p%10)
    if p%10 != 9 {
      slice = append(slice, 10*p + p%10 +1)
    }
  }
  return p
}
