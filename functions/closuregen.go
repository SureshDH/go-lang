package main

import "fmt"

func evenGenrator() func(int) int{
    i := int(0)
    return func(rand int)(ret int){
        ret = i
        i += 2
        fmt.Println(rand)
        return
    }
}

func main(){
    genFunc := evenGenrator()
    fmt.Println(genFunc(22))
    fmt.Println(genFunc(44))
    fmt.Println(genFunc(66))
    fmt.Println(genFunc(88))
    fmt.Println(genFunc(11))
}
