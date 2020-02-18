package main

import (
  "fmt"
  "strconv"
)

func main() {
  var (
    n,a,b,sum,ans int
  )
  fmt.Scanf("%d %d %d", &n, &a, &b)

  for i:=0; i<n+1; i++ {
    sum = describe_sum_of_each_keta(i)
    if sum >= a && sum <= b {
      ans+=i
    }
  }
  fmt.Println(ans)
}

func describe_sum_of_each_keta(x int) int {
  var result int

  xs := strconv.Itoa(x)
  for i:=0; i<len(xs); i++ {
    xi,_ := strconv.Atoi(xs[i:i+1])
    result += xi
  }
  return result
}
