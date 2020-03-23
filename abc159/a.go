package main

import (
  "fmt"
)

func main() {
  var n,m int

  fmt.Scan(&n)
  fmt.Scan(&m)

  fmt.Println(combination(n,2)+combination(m,2))
}

func permutation(n int, k int) int {
    v := 1
    if 0 < k && k <= n {
        for i := 0; i < k; i++ {
            v *= (n - i)
        }
    } else if k > n {
        v = 0
    }
    return v
}

func factorial(n int) int {
    return permutation(n, n - 1)
}

func combination(n int, k int) int {
    return permutation(n, k) / factorial(k)
}
