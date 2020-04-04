package main

import (
	"fmt"
)

func main() {
	var n, m, p int
	var s string
	fmt.Scanf("%d %d", &n, &m)

	acflag := map[int]bool{}
	var ac, wa int
	wabuf := map[int]int{}

	for i := 0; i < m; i++ {
		fmt.Scan(&p)
		fmt.Scan(&s)
		if !acflag[p] {
			switch s {
			case "AC":
				ac++
				acflag[p] = true
				wa += wabuf[p]
			case "WA":
				wabuf[p]++
			}
		}
	}

	fmt.Print(ac)
	fmt.Print(" ")
	fmt.Println(wa)
}
