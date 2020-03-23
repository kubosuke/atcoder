package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var dfs func(s string, mx int)
	dfs = func(s string, mx int) {
		if len(s) == n {
			fmt.Println(s)
			return
		}

		for c := int('a'); c <= mx; c++ {
			nmx := mx
			if c == mx {
				nmx++
			}
			dfs(s+string(c), nmx)
		}
	}
	dfs("", 'a')
}
