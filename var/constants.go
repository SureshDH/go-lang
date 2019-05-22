package main

import "fmt"

const MAX_WT = 514

func main(){

    // you cannot use := for constants
    const constant1 string = "yellow"
    const constant2 = "Green"

    fmt.Println(MAX_WT)
    fmt.Println(constant1)
    fmt.Println(constant2)

    // bool constants
    const published = true
    fmt.Println(published)

    // integer constant
    const number = 515
    var intVar int = number
    var int32Var int32 = number
    var float64Var float64 = number
    var complex64Var complex64 = number
    fmt.Println(number)
    fmt.Println("intVar: ", intVar, "\nint32Var: ", int32Var, "\nfloat64Var: ", float64Var, "\ncomplex64Var: ", complex64Var)
}

// https://dev.to/codehakase/constants-in-go-92o
