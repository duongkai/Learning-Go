package main

import "fmt"

func f (args ...int) []int {
    for gi := 0; gi < len (args); gi++ {
        for gj := gi; gj < len (args); gj++ {
            if args[gi] > args[gj] {
                tmp := args[gi]
                args[gi] = args[gj]
                args[gj] = tmp
            }
        }
    }
    return args
}

func main() {
    fmt.Println (f (8, 4, 19))
}
