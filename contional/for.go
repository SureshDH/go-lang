package main

import "fmt"

func main(){
    i := 0
    for i <= 10 {
        fmt.Println(i)
        i += 1
    }

    fmt.Println("=========================")
    for j := 0; j <= 10; j++{
        fmt.Println(j)
    }

    fmt.Println("=========================")
    for k := 0; k <= 10; k += 1{
        fmt.Println(k)
    }
}
