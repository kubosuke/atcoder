package main

import (
  "fmt"
  // "strconv"
  "math/big"
)

func main() {
  var (
    n,p,ans int64
    s string
  )

  x := new(big.Int)
  z := new(big.Int)
  zero := big.NewInt(0)

  fmt.Scan(&n)
  fmt.Scan(&p)
  fmt.Scan(&s)

  for i:=0; i<len(s); i++ {
    for j:=len(s); j>i; j-- {
      // x, err := strconv.Parseint(s[i:j], 10, 64)
      x, _ := x.SetString(s[i:j], 10)
      P := big.NewInt(p)
      z.Mod(x,P)
      if z == zero {
        ans++
      }
    }
  }
  fmt.Println(ans)
}
