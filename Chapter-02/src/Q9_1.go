package main
import "fmt"

type MyStack struct {
    leng int
    data [10]int
}

func (ms *MyStack) push (n int) {
    defer ms.status()
    if ms.leng + 1 > 10 {
        return
    }
    ms.data[ms.leng] = n
    ms.leng++
}

func (ms *MyStack) pop() int {
    defer ms.status()
    ms.leng--
    return ms.data[ms.leng]
}

func (ms *MyStack) status() {
    fmt.Printf ("%v\n", ms.data)
    fmt.Println ("Current length: ", ms.leng)
}


func main() {
    var s MyStack
    for gi := 10; gi <= 25; gi++ {
        fmt.Println (gi)
        s.push (gi)
    }
}
