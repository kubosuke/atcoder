package main

import "fmt"

func main(){
    var str string
    var ans int = 0
    fmt.Scanf("%s", &str)

    for i:=0; i<3; i++ {
      if str[i:i+1] == "1" {
        ans++
      }
    }

    fmt.Println(ans)
}
