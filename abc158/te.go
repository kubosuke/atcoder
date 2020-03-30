package main

import (
	"fmt"
)

func main() {
var s string = "‚ "
for i := 0; i < len(s); i++ {
	    fmt.Println(s[i]) // e3 81 82
    }
}
