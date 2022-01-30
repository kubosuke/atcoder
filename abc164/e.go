package main

import (
	"fmt"
)

func main() {
	var (
		n,m,s int
	)
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&s)
	u := make([]int, m)
	v := make([]int, m)
	a := make([]int, m)
	b := make([]int, m)
	c := make([]int, n)
	d := make([]int, n)
	for i:=0; i<m; i++ {
		fmt.Scan(&u[i])
		fmt.Scan(&v[i])
		fmt.Scan(&a[i])
		fmt.Scan(&b[i])
	}
	for j:=0; j<n; j++ {
		fmt.Scan(&c[j])
		fmt.Scan(&d[j])
	}
}
