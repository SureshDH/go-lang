package main

import (
    "fmt"
    "sort"
    )

func main(){
    // Note: the last item in the map requires ,(comma)
    players := map[string]int{
        "Saina": 28,
        "Sindhu": 23,
        "Srikanth": 26,
        "Sameer": 24,
        "Malavika": 16,
    }

    pnames := make([]string, 0, len(players))

    for name := range players{
        pnames = append(pnames, name)
    }

    sort.Strings(pnames) // sort the keys alphabetically

    for _, name := range pnames{
        fmt.Println(players[name])
    }
    numbers := []int{28, 38, 49, 1, 10, 5}
    fmt.Println(numbers)
    sort.Ints(numbers) // sort the keys alphabetically
    fmt.Println(numbers)
    flag := sort.IntsAreSorted(numbers) // check whether number sorted or not.
    fmt.Println(flag)

}
