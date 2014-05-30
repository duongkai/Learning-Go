package main

import "fmt"

func prnt (args ...int) {
    for gi, value := range args {
        fmt.Println (gi, " element = ", value)
    }
}

func main() {
    prnt (1, 10, 13, 9, 3)
}
