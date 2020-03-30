package main

import (
  "fmt"
  "strings"
)

func main() {
  str := "1, 1, 1, 2, 1, 2, 1, 5, 2, 2, 1, 5, 1, 2, 1, 14, 1, 5, 1, 5, 2, 2, 1, 15, 2, 2, 5, 4, 1, 4, 1, 51"
  slice := strings.Split(str, ", ")
  var d int
  fmt.Scan(&d)
  fmt.Println(slice[d-1])
}
