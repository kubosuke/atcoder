package main

import (
	"fmt"
	"math"
)

func main() {
	var n,x,top int
	var ans int = 1
	fmt.Scan(&n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&d[i])
	}
	mochinum := len(d)
	top = max(d)

	for j := 0; j < mochinum; j++ {
		x,d = export_max(d)
		if(top > x) {
			ans++
			top = x
		}
	}
	fmt.Println(ans)
}

func export_max(nums []int) (int, []int) {
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

func max(nums []int) int {
    if len(nums) == 0 {
        panic("funciton max() requires at least one argument.")
    }
    res := nums[0]
    for i := 0; i < len(nums); i++ {
        res = int(math.Max(float64(res), float64(nums[i])))
    }
    return res
}
