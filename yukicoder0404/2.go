package main

import (
  "fmt"
)

var n int

func main() {
  var (
    s string
  )
  var p []string
  fmt.Scan(&n)
  fmt.Scan(&s)

  for _, c := range s {
    p = append(p,string(c))
  }

  fmt.Println(task(p))
}

func task(p []string) string {
  var m map[int]int{}
  for i, x := range p {
    if x == "-" {
      if i!=0 {
        if p[i-1] == "o" {
          m[i]++
        }
      }
      if i!=n-1 {
        if p[i+1] == "o" {
          m[i]++
        }
      }
    }
  }
  var max int = -1
  var min int = -1
  for j, y := range m {
    if min!=-1 {
      min=j
    }
    if y==2 {
      return j
    }
    if y==1 {
      if max!=-1 {
        max=j
      }
    }
  }
}
