package main

import (
	"fmt"
)

func main() {
	var (
		s string
	)
	fmt.Scan(&s)

	if (erasedream(s)=="NO" && dreamerase(s) == "NO") {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

func erasedream(str string) string {
	fmt.Println(str)
	if len(str) == 0 {
		return "YES"
	}
	if len(str) > 0 {
		if str[:1] == "r" {
			return erasedream(str[1:])
		}
		if len(str) > 1 {
			if str[:2] == "er" {
				return erasedream(str[2:])
			}
			if len(str) > 4 {
				if str[:5] == "dream" || str[:5] == "erase" {
					return erasedream(str[5:])
				}
				if len(str) > 5 {
					if str[:6] == "eraser" {
						return erasedream(str[6:])
					}
					if len(str) > 6 {
						if str[:7] == "dreamer" {
							return erasedream(str[7:])
						}
					}
				}
			}
		}
	}
	return "NO"
}

func dreamerase(str string) string {
	fmt.Println(str)
	if len(str) == 0 {
		return "YES"
	}
	if len(str) > 6 {
		if str[:7] == "dreamer" {
			return draemerase(str[7:])
		}
		if len(str) > 5 {
			if str[:6] == "eraser" {
				return dreamerase(str[6:])
			}
			if len(str) > 4 {
				if str[:5] == "dream" || str[:5] == "erase" {
					return dreamerase(str[5:])
				}
				if len(str) > 1 {
					if str[:2] == "er" {
						return dreamerase(str[2:])
					}
					if len(str) > 0 {
						if str[:1] == "r" {
							return dreamerase(str[1:])
						}
					}
				}
			}
		}
	}
	return "NO"
}
