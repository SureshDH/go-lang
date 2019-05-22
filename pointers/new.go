package main

import "fmt"

func assginValue(ptr *int){
    *ptr = 515
}

func main(){
    xvar := new(int)
    *xvar = 514
    fmt.Println("Before: ", xvar)
    fmt.Println("Before: ", *xvar)
    assginValue(xvar)
    fmt.Println("After: ", xvar)
    fmt.Println("After: ", *xvar)

}
