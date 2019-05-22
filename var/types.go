package main

import "fmt"

func main(){

    /*
        Go supports Boolean, strings, Numerix(int, float, compplex, byte, rune)
    */

    // Boolean takes on byte
    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(!false)

    // integer types are int8, int16, int32, int64, int
    // uint8, uint16, uint32, uint64, uint
    
    var inumber int = 514
    var unumber uint = 515
    fmt.Println("inumber: ", inumber)
    fmt.Println("unumber: ", unumber)

    // float32, float64
    var fnumber1 float32 = 3.14
    var fnumber2 float32 = 1.414
    fmt.Println("fadd: ", fnumber1+fnumber2)
    fmt.Println("fdiff: ", fnumber1-fnumber2)

    // complex32, complex64
    // if two complex64 gets added then it gets saved in compplex128
    c1 := complex(2, 3)
    c2 := 2 + 47i

    cadd := c1 + c2
    fmt.Println("cadd: ", cadd)

    cmul := c1 * c2
    fmt.Println("cmul: ", cmul)

    // strings
    string1 := "This is a string"
    fmt.Println("string1: ", string1)



}
