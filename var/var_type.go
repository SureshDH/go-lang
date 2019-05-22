package main

import "fmt"

func main(){
    // It doesn't require any variable type if the 
    // value gets assigned at declaration time
    var number = 514
    fmt.Println(number)
    fmt.Println(number)

    // In go you can assigne multiple variables at the same time
    var day, month, year int = 14, 05, 2019
    fmt.Println("This program was written on %d/%d/%d", day, month, year)

    // Multiple assignments with different types are also possible
    var (
        name string = "Yuta Watanade"
        knownfor string = "Deceptive shots"
    )
    fmt.Println(name+" is known for "+knownfor)

    // There is even shorthand for variable decalration
    worldChompian := "Kento Momota"
    fmt.Println("Reigning world chompian is ", worldChompian)

    // Some extra stuff
    firstName, lastName := "Sayaka", "Takahashi"
    fmt.Println("First name :%s, Last Name :%s", firstName, lastName)

    firstName = "Ayaka" // firstName := cannot be used this way
    fmt.Println("First name :%s, Last Name :%s", firstName, lastName)

    // The following type conversion is not possible
    /*
    firstName = 514
    fmt.Println("First name :%s, Last Name :%s", firstName, lastName)
    */

    name, age := "Saina Nehwal", 28
    fmt.Println("Player name: "+name+" and age : %d", age)

}

// https://dev.to/codehakase/types-in-go-ckf
