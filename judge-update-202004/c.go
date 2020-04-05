package main

import (
	"fmt"
)

func main() {
	var a1,a2,a3 int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	s := a1+a2+a3
	fmt.Println((s%
		a1 + s%a2 + s%a3)+1)
}
