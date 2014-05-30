package main
import "fmt"

// define stack
var myStack [10]int
var stackLength int

func push (newElement int) bool {
    defer status()
    if stackLength >= 10 {
        //panic ("length exceeds 10")
        return false
    }
    myStack[stackLength] = newElement
    stackLength++
    return true
}

func pop() int {
    defer status()
    if stackLength == 0 {
        //panic ("length is zero")
        return -1
    }
    stackLength--
    return myStack[stackLength]
}


func status() {
    fmt.Println ("Current stack: ", myStack)
    fmt.Println ("stack length = ", stackLength)
}

func main() {
    status()
    for gi := 10; gi <= 20; gi++ {
        push (gi * 2)
    }
}
