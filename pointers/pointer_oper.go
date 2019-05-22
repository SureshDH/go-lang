package main

import "fmt"

func main(){
    var num int
    num2 := &num
    *num2 = 514
    fmt.Println("pointer num is : ", num)
    fmt.Println("pointer &num is : ", &num)
    fmt.Println("pointer num2 is : ", num2)
    fmt.Println("pointer &num2 is : ", &num2)

    array := []int{2, 5, 8, 1, 6}
    fmt.Println("array: ", array)
    fmt.Println("&array: ", &array)
    fmt.Println("array: ", array[1])
    // POinter arithmetic is not possible in go
    // So following two instructions will compile to fail
    // fmt.Println(num2++)
    // fmt.Println("array++: ", (array++)*)
}
