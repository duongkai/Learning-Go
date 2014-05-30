package main
import "fmt"

func main() {
    a := []int{1, 3, 4}
    fmt.Println (simple_map (a, square))
}

func square (x int) int {
    return x * x
}

func simple_map (args []int, f func(int) int) []int {
    for gi, value := range args {
        args[gi] = f (value)
    }
    return args
}
