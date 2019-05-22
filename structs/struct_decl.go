package main

import (
    "fmt"
    )

type Person struct{
    name string
    age int
}


func main(){
    fmt.Println(Person{"Saina Nehwal", 28})
    fmt.Println(Person{name: "Sindhu PV", age: 23})
    fmt.Println(Person{name: "Srikanth Kidambi"})
    fmt.Println(&Person{name: "Sameer Verma", age: 28})
    fmt.Println(&Person{"Naomi Osaka", 21})
    fmt.Println(&Person{name: "Malavika Bansod"})
    fmt.Println(&Person{"Lakshya Sen", 19})
}
