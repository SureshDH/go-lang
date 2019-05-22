package main

import "fmt"

func last(){
    fmt.Println("From function last")
}

func first(){
    fmt.Println("From the function first")
}

func main(){
    /*
        Here the defer function gets called after the
        first function.
        Defer is used to cleanup the resources like file close, etc.
        Output: 
            From the function first
            From function last
    */
    defer last()
    first()
}
