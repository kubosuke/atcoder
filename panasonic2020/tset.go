package main

import (
  "fmt"
  "math/big"
)

func main() {
  bigTwo := big.NewInt(2)
  bigSixteen := big.NewInt(16)

  sqrt := bigTwo.Sqrt(bigSixteen)
  fmt.Println(sqrt)
}
