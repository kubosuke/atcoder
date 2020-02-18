package main

import (
	"fmt"
)

func main() {
	var str string = "abcdefghijklmnopqrstuvwxyz"
	var si string

  fmt.Scanf("%s", &si)

	for i := 0; i < len(str)-1; i++ {
		if si == str[i:i+1] {
			fmt.Println(str[i+1:i+2])
		}
	}
}
