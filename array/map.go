package main

import "fmt"

func main(){

    // The following won't work becase we assigning
    // a value to an nil address
    /*
        var x map[string]int
        x["hello"] = 514
        fmt.Println(x["hello"])
    */

    // Different ways to declare maps
	/*
        elements := map[string]string{
		  "H":  "Hydrogen",
		  "He": "Helium",
		  "Li": "Lithium",
		  "Be": "Beryllium",
		  "B":  "Boron",
		  "C":  "Carbon",
		  "N":  "Nitrogen",
		  "O":  "Oxygen",
		  "F":  "Fluorine",
		  "Ne": "Neon",
		}
	*/
	elements := map[string]map[string]string{
		"H": map[string]string{
		  "name":"Hydrogen",
		  "state":"gas",
		},
		"He": map[string]string{
		  "name":"Helium",
		  "state":"gas",
		},
		"Li": map[string]string{
		  "name":"Lithium",
		  "state":"solid",
		},
		"Be": map[string]string{
		  "name":"Beryllium",
		  "state":"solid",
		},
		"B":  map[string]string{
		  "name":"Boron",
		  "state":"solid",
		},
		"C":  map[string]string{
		  "name":"Carbon",
		  "state":"solid",
		},
		"N":  map[string]string{
		  "name":"Nitrogen",
		  "state":"gas",
		},
		"O":  map[string]string{
		  "name":"Oxygen",
		  "state":"gas",
		},
		"F":  map[string]string{
		  "name":"Fluorine",
		  "state":"gas",
		},
		"Ne":  map[string]string{
		  "name":"Neon",
		  "state":"gas",
		},
	  }

	  if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	  }
    // The following should work
    x := make(map[string]int)
    x["Sureh"] = 26
    x["Ramesh"] = 29
    fmt.Println(x["Sureh"], x["Ramesh"])

    // Key can an integer also
    y := make(map[int]int)
    y[1] = 514
    y[2] = 515
    y[3] = 562
    fmt.Println(y[1], y[2], y[3])
    delete(y, 3)
    fmt.Println(y[1], y[2], y[3])

    if age, ok := x["Suresh"]; ok {
        fmt.Println(age, ok)
    } else {
        fmt.Println(age, ok)
    }
}
