package main

import (
  "fmt"
)

func main() {
  var (
    n,a,b,i int
    ans int
  )
  fmt.Scan(&n)
  fmt.Scan(&a)
  fmt.Scan(&b)

  for i=1; i<n+1; i++ {
    if i==a || i==b {
      continue
    } else {
      ans += combination(n,i)
    }
  }
  fmt.Println(ans)
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
