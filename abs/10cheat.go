package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		s string
	)
	fmt.Scan(&s)

	fmt.Println(erasedream(s))
}

func erasedream(str string) string {
	str = strings.Replace(str, "erase", "E", -1)
	str = strings.Replace(str, "dream", "D", -1)
	str = strings.Replace(str, "Der", "", -1)
	str = strings.Replace(str, "Er", "", -1)
	str = strings.Replace(str, "D", "", -1)
	str = strings.Replace(str, "E", "", -1)
	if (len(str) == 0) {
		return "YES"
	} else {
		return "NO"
	}
}
