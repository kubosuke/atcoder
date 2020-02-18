package main

import (
	"fmt"
	"math"
)

func main() {
	var n, drow, alice, bob int
	fmt.Scan(&n)
	a := make([]int, n)
  for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	round := len(a)
	for i := 0; i < round; i++ {
		drow, a = max(a)
		switch i % 2 {
		case 0:
			alice += drow
		case 1:
			bob += drow
		}
	}
	fmt.Println(alice - bob)
}

func max(nums []int) (int, []int) {
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
