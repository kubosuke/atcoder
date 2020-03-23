package main

import (
	"fmt"
)

func main() {
	var x, y, h, i, buf float64
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&h)
	if x < y {
		buf, i = task(1000*x, h, i)
		_, i = task(1000*y, buf, i)
	} else {
		buf, i = task(1000*y, h, i)
		_, i = task(1000*x, buf, i)
	}
	fmt.Println(i)
}

func task(a, h, i float64) (float64, float64) {
	for {
		if a > h {
			a /= 2
			h *= 2
			i++
		}
		if a <= h {
			return h, i
		}
	}
}
