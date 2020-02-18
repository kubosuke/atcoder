package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
  "strconv"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	stdin.Scan()
	nkm := strings.Split(stdin.Text(), " ")
  n,_ := strconv.Atoi(nkm[0])
  k,_ := strconv.Atoi(nkm[1])
  m,_ := strconv.Atoi(nkm[2])

	stdin.Scan()
	a := strings.Split(stdin.Text(), " ")

  var present_score int = 0
  for i:=0; i<len(a); i++ {
    buf,_ := strconv.Atoi(a[i])
    present_score += buf
  }

  need_score := n*m - present_score

  switch {
  case need_score <= 0:
    fmt.Println("0")
  case need_score > k:
    fmt.Println("-1")
  default:
    fmt.Println(need_score)
  }
}
