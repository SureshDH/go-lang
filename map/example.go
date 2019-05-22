package main

import "fmt"

func main(){
    // Note: the last item in the map requires ,(comma)
    players := map[string]int{
        "Saina": 28,
        "Sindhu": 23,
        "Srikanth": 26,
        "Sameer": 24,
        "Malavika": 16,
    }

    // Each print will print different ouput
    fmt.Println(players)
    fmt.Println(players)
    fmt.Println(players)
    fmt.Println(players)
    fmt.Println(players)
    fmt.Println(players)


    // Other way to iterate over the map
    for name, age := range players{
        fmt.Println(name+": ", age)
    }
}
