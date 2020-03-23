package main

import (
 "fmt"
 "math"
)

// 引数の関数の x1 から x2 までの定積分の値の近似値を台形公式により求めて返す関数
// (エラーチェックなし)
// https://ja.wikipedia.org/wiki/%E5%8F%B0%E5%BD%A2%E5%85%AC%E5%BC%8F
const N_SUBINTERVAL = 4000 // 大きくするほど積分の精度がよくなる
var a,b float64

func integrate(f func(float64) float64, x1, x2 float64) float64 {
 if x1 == x2 {
 return 0
 }

 step := ((x2 - x1) / N_SUBINTERVAL)
 sum := 0.5 * f(x1)
 x := x1
 for i := 1; i < N_SUBINTERVAL; i++ {
 x += step
 sum += f(x)
 }
 sum += 0.5 * f(x2)

 return sum * step
}

// f(x) = 4 * (1-x^2)^(1/2)
func testfunc(x float64) float64 {
 return (x-a) * (x-b)
 return 4 * math.Sqrt(1-math.Pow(x, 2))
}

func main() {
  fmt.Scan(&a)
  fmt.Scan(&b)
 // answer equals pi.
 fmt.Println(math.Abs(integrate(testfunc, a, b)))
}
