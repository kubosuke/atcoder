package main

import (
	"fmt"
	"sort"
)

var n int

func main() {
	var x, y, z int
	fmt.Scan(&n)
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&z)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Sort(sort.IntSlice(a))
	fmt.Println(task(a, x, y, z))
}

func task(a []int, x, y, z int) string {
	for a[len(a)-1] > -1 {
    fmt.Println(a)
		if x+y+z == 0 {
			return "No"
		}
		if x != 0 {
			x--
			a[0] -= 1000
			if a[0] < 0 {
				a = a[1:]
			}
			continue
		}
		if y != 0 {
			y--
			a[0] -= 5000
			if a[0] < 0 {
				a = a[1:]
			}
			continue
		}
		if z != 0 {
			z--
			a[0] -= 10000
			if a[0] < 0 {
				a = a[1:]
			}
			continue
		}
	}
	if a[len(a)-1] > -1 {
		return "No"
	} else {
		return "Yes"
	}
}
