package main

import "fmt"

func factorial(x uint) uint{
    if x == 0 || x == 1{
        return 1
    }
    return x * factorial(x-1)
}

func main(){
    var number uint = 5
    fmt.Println(factorial(number))
}
