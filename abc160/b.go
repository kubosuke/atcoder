package main

import (
  "fmt"
)

func main() {
  var x,ans int
  fmt.Scan(&x)

  for x>499 {
    x -= 500
    ans += 1000
  }

  for x > 4 {
    x -= 5
    ans += 5
  }
  fmt.Println(ans)

}
