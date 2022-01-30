package main

import (
  "fmt"
)

func main() {
  var (
    a,b,c,d int
  )

  fmt.Scan(&a)
  fmt.Scan(&b)
  fmt.Scan(&c)
  fmt.Scan(&d)

  fmt.Println(task(a,b,c,d))

}

func task(a,b,c,d int) string {
  for {
    c -= b
    if c<=0 {
      return "Yes"
    }
    a -= d
    if a<=0 {
      return "No"
    }
  }
}
