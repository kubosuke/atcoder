package main

import (
	"fmt"
)

func main() {
	var n,m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	s := make([]int, m)
	c := make([]int, m)

	for i:=0; i<m; i++ {
		fmt.Scan(&s[i])
		fmt.Scan(&c[i])
	}

	fmt.Println(check())
}

func check(n int, keta []int, numlist []int) {
	
}
