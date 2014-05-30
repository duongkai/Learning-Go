package main

import "fmt"

func main() {
    for gi := 1; gi < 10; gi++ {
        fmt.Println (fibonacci (gi))
    }
    fmt.Println (fibonaccis (20))
}

func fibonacci (n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    return fibonacci (n-1) + fibonacci (n-2)
}

func fibonaccis (num int) []int {
    fibos := make ([]int, num)
    fibos[0] = 1
    fibos[1] = 1
    for gi := 2; gi < num; gi++ {
        fibos[gi] = fibos[gi - 1] + fibos[gi - 2]
    }
    return fibos
}
