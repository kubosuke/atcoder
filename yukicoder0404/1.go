package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getScanner(fp *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 1000005), 1000005)
	return scanner
}

func getNextString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func getNextInt(scanner *bufio.Scanner) int {
	i, _ := strconv.Atoi(getNextString(scanner))
	return i
}

func getNextInt64(scanner *bufio.Scanner) int64 {
	i, _ := strconv.ParseInt(getNextString(scanner), 10, 64)
	return i
}

func getNextUint64(scanner *bufio.Scanner) uint64 {
	i, _ := strconv.ParseUint(getNextString(scanner), 10, 64)
	return i
}

func getNextFloat64(scanner *bufio.Scanner) float64 {
	i, _ := strconv.ParseFloat(getNextString(scanner), 64)
	return i
}

func main() {
	fp := os.Stdin
	wfp := os.Stdout
	cnt := 0

	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)
	solve(scanner, writer)
	for i := 0; i < cnt; i++ {
		fmt.Fprintln(writer, "-----------------------------------")
		solve(scanner, writer)
	}
	writer.Flush()
}

func solve(scanner *bufio.Scanner, writer *bufio.Writer) {
	n := getNextInt(scanner)
	x := getNextInt(scanner)
	y := getNextInt(scanner)
	z := getNextInt(scanner)

	aa := make([]int, n)
	for i := 0; i < n; i++ {
		aa[i] = getNextInt(scanner)
	}

	loop := [3]int{
		z,
		y,
		x,
	}
	money := [3]int{
		10000,
		5000,
		1000,
	}
	for l := 0; l < 3; l++ {
		sort.Ints(aa)
		for i := n - 1; i >= 0; i-- {
			if aa[i] < 0 {
				continue
			}

			if loop[l] > 0 {
				used := aa[i] / money[l]
				if loop[l] >= used {
					aa[i] -= money[l] * used
					loop[l] -= used
				}
			}
		}
		sort.Ints(aa)
		for i := n - 1; i >= 0; i-- {
			if aa[i] < 0 {
				continue
			}

			if loop[l] > 0 {
				aa[i] -= money[l]
				loop[l]--
			}
		}
	}
	if complete(aa) {
		fmt.Fprintln(writer, "Yes")
		return
	}
	fmt.Fprintln(writer, "No")
}

func complete(aa []int) bool {
	n := len(aa)
	for i := 0; i < n; i++ {
		if aa[i] >= 0 {
			return false
		}
	}
	return true
}
