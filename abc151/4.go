package main

import (
	"fmt"
)

func main() {
	var h,w int
	fmt.Scanf("%d %d", &h, &w)

	var tmp string
	var s []string
	for i:=0; i<h; i++ {
		fmt.Scan(&tmp)
		for _, c := range tmp {
			s = append(s, string(c))
		}
	}
	fmt.Println(s)
}
