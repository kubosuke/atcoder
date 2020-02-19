package main

import (
  "fmt"
)

func main() {
  var (
    s,t,u string
    a,b int
  )
  fmt.Scan(&s)
  fmt.Scan(&t)
  fmt.Scan(&a)
  fmt.Scan(&b)
  fmt.Scan(&u)

  switch(u) {
  case s:
    a--
  case t:
    b--
  }
  fmt.Print(a)
  fmt.Print(" ")
  fmt.Println(b)
}
