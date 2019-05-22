package main

import "fmt"

func main(){
    /*
        The following example doesn't work because we are
        not defering the recover function.
            panic("PANIC")
            str := recover()
            fmt.Println(str)
        Also the defer function should always above the panic
        function else it will simply abort without any message
    */
    defer func(){
        str := recover()
        fmt.Println(str)
    }()
    panic("PANIC")
}
