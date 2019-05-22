package main

import "fmt"

func variadic(args ...int){
    for _, value := range args{
        fmt.Println(value)
    }
}

func main(){
    variadic(12, 14, 23, 24, 25)
}
