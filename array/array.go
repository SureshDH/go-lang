package main

import "fmt"

func main(){

    // By default all the variable of an integer array
    // are initialized to zero
    // string array to null space
    // float array to zero
    var x [5]int
    x[4] = 514
    fmt.Println(x)

    // different ways of decaring array
	/*
		x := [5]float64{ 98, 93, 77, 82, 83 }

		x := [5]float64{
		  98,
		  93,
		  77,
		  82,
		  83,
		}

		x := [4]float64{
		  98,
		  93,
		  77,
		  82,
		  // 83,
		}
	*/

    // initializing an array
    // var arr1 = [5]int{1, 2, 3, 4, 5}
    var arr1 = []int{1, 2, 3, 4, 5}
    fmt.Println(arr1)

    var total float32
    for i := 0; i < len(arr1); i += 1 {
        total += float32(arr1[i])
    }
    fmt.Println(total/float32(len(arr1)))

    total = 0
    fmt.Println("#####################################")
    // different way of retriving value from array
    for i, value := range arr1 {
        total += float32(value)
        fmt.Println(i)
    }
    fmt.Println(total/float32(len(arr1)))

    total = 0
    fmt.Println("#####################################")
    // avoid the use of unused variables
    for _, value := range arr1 {
        total += float32(value)
    }
    fmt.Println(total/float32(len(arr1)))

    // array as function parameter
    avg := calculateAvg(arr1)
    fmt.Println("Avg: ", avg)

}

func calculateAvg(arr []int) float64{
    total := 0.0
    for _, value := range arr {
        total += float64(value)
    }
    return total/float64(len(arr))

}
