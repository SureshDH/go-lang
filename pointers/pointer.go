package main

import "fmt"

func changeValue(x *int){
    *x = 514
}

func main(){
    x := 10
    fmt.Println(x)
    changeValue(&x)
    fmt.Println(x)
}
