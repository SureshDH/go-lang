package main

import (
    "fmt"
    "strings"
    )

func main(){
    // str := "Warangal"
    fmt.Println(strings.HasPrefix("Warangal", "war"))
    fmt.Println(strings.HasPrefix("Warangal", "War"))
}
