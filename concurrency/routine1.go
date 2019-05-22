package main

import "fmt"

func f(number int){
    for i:= 0; i<=10; i += 1{
        fmt.Println(number, ":", i)
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
