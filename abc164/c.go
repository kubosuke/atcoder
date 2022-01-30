package main

import (
  "fmt"
)

func main() {
  var (
    n int
  )
  fmt.Scan(&n)
  s := make([]string, n)
  for i:=0; i<n; i++ {
    fmt.Scan(&s[i])
  }

  s2 := removeDuplicate(s)
  if len(s2) > n {
    fmt.Println(n)
  } else {
    fmt.Println(len(s2))
  }
}

func removeDuplicate(args []string) []string {
    results := make([]string, 0, len(args))
    encountered := map[string]bool{}
    for i := 0; i < len(args); i++ {
        if !encountered[args[i]] {
            encountered[args[i]] = true
            results = append(results, args[i])
        }
    }
    return results
}
