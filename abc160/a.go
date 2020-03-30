package main

import (
  "fmt"
)

func main() {
  var str string
  fmt.Scan(&str)

  if str[2:3] == str[3:4] && str[4:5] == str[5:6] {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
}
