package main

import (
	"fmt"
)

func main() {
	var h, w, ans int64
	fmt.Scan(&h)
	fmt.Scan(&w)

	if h == 1 || w == 1 {
		ans = 1
	} else {
		if h%2 == 0 || w%2 == 0 {
			ans = (h * w) / 2
		} else {
			ans = 1 + ((h * w) / 2)
		}
	}

	fmt.Println(ans)
}
