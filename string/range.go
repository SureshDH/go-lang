package main

import (
    "fmt"
    )

func main(){
    for k, v := range "abcdefgzABCDEFGZ"{
        fmt.Println(k, v)
    }
}
