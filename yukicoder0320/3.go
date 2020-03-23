package main

import (
  "fmt"
)

var n,d,k,ans int

func main() {
  fmt.Scan(&n)
  fmt.Scan(&d)
  fmt.Scan(&k)

  task(0,0)

  fmt.Println(ans)
}

func task(counter,sum int) {
  if counter == n {
    if sum == k {
      ans++
    }
  } else {
    if sum < k {
      for j:=1; j<d+1; j++ {
        task(counter+1, sum+j)
      }
    }
  }
}
