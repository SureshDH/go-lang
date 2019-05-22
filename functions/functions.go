package main

import "fmt"

// Function without parameters
// Function starts with func keyword
func hello(){
    fmt.Println("Hello There!")
}

// Function with parameters
func helloArg(name string){
    fmt.Println("Hello "+name);
}

// Function returning variables
func funcRet(number int)(int, int){
    return 514+number, 514-number
}

// If all params are of same type then it cane be
// mentioned in the end of the list of args
func nargs(number1, number2, number3 int) int{
    sum := number1 + number2 + number3
    return sum
}

// If return variable name is mention in the function
// defination then a simple return statment is enough
func retExample(number1, number2 int) (sum int){
    sum = number1 + number2
    return
}

func main(){
    hello()
    helloArg("Desktop")
    fmt.Println(funcRet(514))
    fmt.Println(nargs(2, 3, 4))
    fmt.Println(retExample(2, 3))
}

// Ref: https://dev.to/codehakase/functions-in-go-3dd
