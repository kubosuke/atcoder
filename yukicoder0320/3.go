package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := readInt()
	d := readInt()
	K := readInt()

	t0 := make([]int, d*N+K+1)
	t0[0] = 1
	for i := 0; i < N; i++ {
		t1 := make([]int, d*N+K+1)
		t := 0
		for j := 0; j < d; j++ {
			t += t0[j]
			t %= 1000000007
			t1[j] = t
		}
		for j := d; j < (i+1)*d; j++ {
			t -= t0[j-d]
			if t < 0 {
				t += 1000000007
			}
			t += t0[j]
			t %= 1000000007
			t1[j] = t
		}
		t0 = t1
	}
	fmt.Println(t0[K-N])
}

const (
	ioBufferSize = 1 * 1024 * 1024 // 1 MB
)

var stdinScanner = func() *bufio.Scanner {
	result := bufio.NewScanner(os.Stdin)
	result.Buffer(make([]byte, ioBufferSize), ioBufferSize)
	result.Split(bufio.ScanWords)
	return result
}()

func readString() string {
	stdinScanner.Scan()
	return stdinScanner.Text()
}

func readInt() int {
	result, err := strconv.Atoi(readString())
	if err != nil {
		panic(err)
	}
	return result
}
