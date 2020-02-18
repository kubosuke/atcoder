package main

import (
	"fmt"
)

func main() {
	var (
		n,y int
		a,b,c int = 0,0,0
	)
	fmt.Scan(&n)
	fmt.Scan(&y)

	for m:=0; m<n+1; m++ {
		for g:=0; g<n+1; g++ {
			s := n-m-g
			if s<0 {
				s=0
			}
				if m*10000 + g*5000 + s*1000 == y && m+g+s==n {
					a=m
					b=g
					c=s
					goto yeah
			}
		}
	}

	if (a+b+c==0) {
		a,b,c=-1,-1,-1
	}

yeah:
	fmt.Print(a)
	fmt.Print(" ")
	fmt.Print(b)
	fmt.Print(" ")
	fmt.Println(c)
}
