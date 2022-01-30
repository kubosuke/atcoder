package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	task(n)
}

func task(x int) {
	for i := -118; i < 120; i++ {
		for j := -119; j < 119; j++ {
			if (i*i*i*i*i)-(j*j*j*j*j) == x {
				fmt.Print(i)
				fmt.Print(" ")
				fmt.Println(j)
				return
			}
		}
	}
}
