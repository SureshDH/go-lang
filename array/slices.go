package main

import "fmt"

func main(){
	x := [10]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println(x)

    // make function is used to create slices
    slice1 := make([]float32, 5)
    fmt.Println(slice1)

    // 10 represents capacity of the underlying array
    slice2 := make([]float32, 5, 10)
    fmt.Println(slice2)

    // by using [low: high]
    slice3 := x[2:8]
    fmt.Println(slice3)

    // differenet slicing ways
    fmt.Println(x[3:])
    fmt.Println(x[:7])
    fmt.Println(x[:])
    fmt.Println(x[1:len(x)])

    // can append slices

    slice4 := append(slice3, 514, 515, 516)
    fmt.Println(slice4)
}
