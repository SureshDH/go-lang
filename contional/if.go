package main

import "fmt"

func main(){

    // The else keyword need to right beside the closing
    // brace of the if block else go throws an error.
    for i := 1; i <= 10; i += 1{
        if i%2 == 0 {
            fmt.Println("Divisible by 2")
        } else if  i%3 == 0{
            fmt.Println("Divisable by 3")
        } else if  i%4 == 0{
            fmt.Println("Divisable by 4")
        }
    }
}
