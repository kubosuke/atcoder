package main

import (
  "fmt"
)

func main() {
  var s string = "3543"
  for i:=0; i<len(s); i++ {
    for j:=len(s); j>i; j-- {
      fmt.Println(s[i:j])
    }
  }
}
