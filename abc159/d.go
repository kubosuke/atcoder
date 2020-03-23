// https://detail.chiebukuro.yahoo.co.jp/qa/question_detail/q1385775538

package main

import (
	"fmt"
)

func main() {
	var n,x int
  var a,b []int
  var res int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
    a = append(a,x)
	}

  for j:=0; j<n; j++ {
    res = 0
    b = delete(a, j)
    for len(b)>0 {
      b, x = remove(b, b[0])
      res += combination(x,2)
    }
    fmt.Println(res)
  }
}

func delete(s []int, index int) []int {
    ret := make([]int, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}

func remove(ints []int, search int) ([]int, int) {
    result := []int{}
    res := 0
    for _, v := range ints {
        if v != search {
            result = append(result, v)
        } else {
          res++
        }
    }
    return result, res
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
	return permutation(n, n-1)
}

func combination(n int, k int) int {
	return permutation(n, k) / factorial(k)
}
