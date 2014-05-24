package main

import "fmt"
//import "math/big"

func main() {
    sum := 0
    for gi := 0; gi <= 10; gi++ {
        sum += gi
    }
    fmt.Println ("sum = ", sum)
    if sum >= 40 {
        fmt.Println ("Salsa")
    } else {
        fmt.Println ("Poly")
    }
    var arr [4]int
    fmt.Println (arr)
    fmt.Println ("Length of the array = ", len (arr))
    for gi := 0; gi < len (arr); gi++ {
        arr[gi] = gi * gi
    }
    fmt.Println (arr)
    // slice define
    a1 := []int{1, 2, 3, 10}
    a2 := []int{4, 6, 9}
    n1 := copy (a1, a2)
    fmt.Println (n1, a1)

    // map
    monthdays := map[string]int{
        "Jan": 31, "Feb": 28, "Mar": 31,
    }
    fmt.Println (monthdays["Feb"])
    for month, days := range monthdays {
        fmt.Println (month, days)
    }
    for _, ge := range (a1) {
        fmt.Println (ge)
    }
}
