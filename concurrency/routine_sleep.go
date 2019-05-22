package main

import (
    "fmt"
    "time"
    "math/rand"
    )

func f(number int){
    for i:= 0; i<=10; i += 1{
        fmt.Println(number, ":", i)
        secs := time.Duration(rand.Intn(250))
        time.Sleep(time.Millisecond * secs)
    }
}

func main(){
    for i := 0; i <= 10; i += 1{
        go f(i)
    }
    var input string
    fmt.Scanln(&input)
    fmt.Println(input)
}
