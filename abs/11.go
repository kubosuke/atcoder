package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	t := make([]int, n)
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
		fmt.Scan(&x[i])
		fmt.Scan(&y[i])
	}

	fmt.Println(is_travelable(t, x, y))
}

func is_travelable(t, x, y []int) string {
	t = append([]int{0}, t...)
	x = append([]int{0}, x...)
	y = append([]int{0}, y...)
	for i := 0; i < len(t)-1; i++ {
		time_difference := t[i+1] - t[i]
		distance := int(math.Abs(float64(x[i+1]-x[i])))+int(math.Abs(float64(y[i+1]-y[i])))
		if time_difference < distance {
			return "No"
		} else {
			if (time_difference-distance)%2 != 0 {
				return "No"
			}
		}
	}
	return "Yes"
}
