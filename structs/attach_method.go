package main

import "fmt"

type Number int

func (i *Number) Add(add Number){
    *i += add
}

func main(){
    // Declaring like below is not going to recomended.
    // for built-in types it has to be declared explicilty
    // i := Number
    var i Number
    i = 514
    i.Add(1)
    fmt.Println(i)
}
