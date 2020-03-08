package main

import "fmt"

var mod = int64(1e9) + 7

func main() {
	var n, a, b int64
	fmt.Scanf("%d %d %d\n", &n, &a, &b)

	nways := power(2, n) - nChooseK(n, a) - nChooseK(n, b) - 1
	for nways < 0 {
		nways += mod
	}

	fmt.Printf("%d\n", nways%mod)
}

func nChooseK(n, k int64) int64 {
	ret := int64(1)
	f := int64(1)
	for i := int64(1); i <= k; i++ {
		ret = (ret * (n - i + 1)) % mod
		f = (f * i) % mod
	}
	return (ret * power(f, mod-2)) % mod
}

func power(a, b int64) int64 {
	ret := int64(1)
	for b > 0 {
		if b%2 != 0 {
			ret = (ret * a) % mod
		}
		b /= 2
		a = (a * a) % mod
	}
	return ret
}
