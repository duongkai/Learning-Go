package main

import "fmt"

func Average (a []float64) float64 {
    n := len (a)
    sum := 0.0
    for _, gi := range a {
        sum += gi
    } 
    return sum / float64 (n)
}

func main() {
    a := []float64{1.0, 2.0, 3.0}
    fmt.Println (Average (a))
}