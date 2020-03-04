package main

import (
	"fmt"
)

func main() {
	var card [3][3]int
	var card_b [3][3]bool
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&card[i][j])
		}
	}
	var n int
	fmt.Scan(&n)
	b := make([]int, n)
	for j := 0; j < n; j++ {
		fmt.Scan(&b[j])
	}

	for _, num := range b {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if card[i][j] == num {
					card_b[i][j] = true
				}
			}
		}
	}

	fmt.Println(check(card_b))
}

func check(c [3][3]bool) string {
	for i := 0; i < 3; i++ {
		if c[i][0] && c[i][1] && c[i][2] || c[0][i] && c[1][i] && c[2][i] {
			return "Yes"
		}
	}
	if c[0][0] && c[1][1] && c[2][2] || c[0][2] && c[1][1] && c[2][0] {
		return "Yes"
	}
	return "No"
}
