package main

import "fmt"

type Dog struct{
    Animal
}

type Animal struct{
    Age int
}

func (a *Animal) Move(){
    fmt.Println("Animal Moved")
}

func (a *Animal) getAge(){
    fmt.Println("Age of animal is", a.Age)
}


func main(){
    d := Dog{}
    d.Age = 15
    d.Move()
    d.getAge()
}
