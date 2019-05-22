package main

import (
    "fmt"
    "bytes"
    )

func main(){
    fmt.Println(bytes.EqualFold([]byte("Go"), []byte("go")))
}
