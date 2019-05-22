package main

import (
    "fmt"
    "/home/sdhara/go/src/avg/mathh"
    )

func main(){
    marks := []float32{88, 86, 80, 72, 88, 89}
    avg := mathh.Avarage(marks)
    fmt.Println(avg)
}
