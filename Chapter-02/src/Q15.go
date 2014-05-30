package main
import "fmt"

func main() {
//    p := plusTwo
//    fmt.Println (p (10))
    p := plusTwoX()
    fmt.Println (p (10))
    p_1 := plusTwoX_1 (10)
    fmt.Println (p_1 (10))
}

func plusTwo (n int) int {
    return n + 2
}

func plusTwoX() func(int) int {
    return func (x int) int { return x + 2 }
}

func plusTwoX_1 (x int) func(int) int {
    return func (y int) int { return x + y }
}
