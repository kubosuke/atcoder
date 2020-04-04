package main

import (
	"fmt"
)

func main() {
	var n, x, y, z int
	fmt.Scan(&n)
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&z)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(task(a, x, y, z))
}

func task(a []int, x, y, z int) string {
    p,i := Max(a)
    if p < 0 || x+y+z == 0 {
      return "No"
    }
		for z!=0 {
			z--
			a[i] -= 10000
      p,i = Max(a)
		}
		for y != 0 {
			y--
			a[i] -= 5000
      p,i = Max(a)
		}
		if x != 0 {
			x--
			a[i] -= 1000
      p,i = Max(a)
		}
    p,_ = Max(a)
	if p>-1 {
		return "No"
	} else {
		return "Yes"
	}
}

func task2(a []int) bool {
	for _, n := range a {
		if n > -1 {
			return true
		}
	}
	return false
}

func Max(array []int) (int, int) {
  var key,value,i int
	var max int = array[0]
	for i, value = range array {
		if max < value {
			max = value
			key = i
		}
	}
	return max, key
}
