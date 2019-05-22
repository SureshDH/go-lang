package main

import (
    "fmt"
    "time"
    )

// func producer2(c chan<- string){
// Above example is used only for sending the traffic
// in one direction. IF don't specify that then it acts
// bi-directional channel
// bi-directional channel can be passed to a function that
// takes send-only or receive-only channels, but the reverse is not true
func producer2(c chan string){
    for i:=0; ; i++ {
        c<-"Hello "
    }
}

func producer1(c chan string){
    for i:=0; ; i++ {
        c<-"World!"
    }
}

// func consumer(c <-chan string){
func consumer(c chan string){
    for {
        msg := <-c
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
   }
}

func main(){
    var c chan string = make(chan string)

    go producer1(c)
    go producer2(c)
    go consumer(c)

    var input string
    fmt.Scanln(&input)
}

