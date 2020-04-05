package main

import (
	"fmt"
  "sort"
)

var k, c int

func main() {
	var (
		n   int
		str string
		d   []int
	)
	fmt.Scan(&n)
	fmt.Scan(&k)
	fmt.Scan(&c)
	fmt.Scan(&str)

	for i, s := range str {
		if string(s) != "x" {
			d = append(d, i+1)
		}
	}
  ans := same(extract(All(d)))
  sort.Sort(sort.IntSlice(ans))
  for _, a := range ans {
    fmt.Println(a)
  }
}

func All(set []int) (subsets [][]int) {
	length := uint(len(set))

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []int

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		if len(subset) == k {
			// add subset to subsets
			subsets = append(subsets, subset)
		}
	}
	return subsets
}

func extract(sets [][]int) [][]int {
	var res [][]int
	for _, set := range sets {
		if task(set) {
			res = append(res, set)
		}
	}
  return res
}

func task(set []int) bool {
	for i := 0; i < len(set)-1; i++ {
		if set[i+1]-set[i] < c+1 {
			return false
		}
	}
  return true
}

func same(sets [][]int) []int {
  var res []int
  ma := map[int]int{}

  for _, set := range sets {
    for _, se := range set {
      ma[se]++
    }
  }

  for i, m := range ma {
    if m==len(sets) {
      res = append(res, i)
    }
  }
  return res
}
