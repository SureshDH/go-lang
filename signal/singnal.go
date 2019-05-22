package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    )

func main(){
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    signal.Notify(sigs,  syscall.SIGTERM, syscall.SIGINT)

    go func(){
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    fmt.Println("Waiting for the signal")
    <-done
    fmt.Println("exiting")
}
