package main

import (
  "fmt"
  "math"
)

func main() {
  var (
    n,k int
    buf float64
  )
  fmt.Scan(&n)
  fmt.Scan(&k)
  p := make([]float64, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&p[i])
  }

  choose_dices := make([]float64, k)
  for j:=0; j<k; j++ {
    choose_dices[j],p = export_max(p)
  }

  for l:=0; l<len(choose_dices); l++ {
    buf += expected_value(choose_dices[l])
  }

  fmt.Println(buf)
}

func export_max(nums []float64) (float64, []float64) {
	var j int

	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		if res < int(math.Max(float64(res), float64(nums[i]))) {
			res = int(math.Max(float64(res), float64(nums[i])))
			j = i
		}
	}
	return res, append(nums[:j],nums[j+1:]...)
}

func expected_value(x float64) float64 {
  var buf float64
  for i:=1; i<int(x)+1; i++ {
    buf += float64(i)
  }
  fmt.Println("buf and x is: ", buf, x)
  return buf/x
}
