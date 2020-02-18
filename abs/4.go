package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

  fmt.Println(count_devide_2(n,nums))

}

func count_devide_2(n int, targets []int) int {
  var ans int

  for {
		for i, target := range targets {
			if target%2 == 0 {
				targets[i] /= 2
			} else {
				return ans
			}
		}
		ans++
	}
}
