package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	n,k := nextInt(), nextInt()
	a := nextInts(n)
	fmt.Println(a,n,k)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextInts(n int) []int {
	sl := make([]int, n)
	for i := range sl {
		sl[i] = nextInt()
	}
	return sl
}
