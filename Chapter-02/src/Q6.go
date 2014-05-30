package main

import "fmt"

func avg (a []float64) float64 {
    if len (a) == 0 {
        return 0;
    }
    sum := 0.0
    for gi := 0; gi < len (a); gi++ {
        sum += a[gi]
    }
    return sum / float64(len (a))
}

func main() {
    a := []float64{1, 3.2, 4.5}
    fmt.Println (avg (a))
}
