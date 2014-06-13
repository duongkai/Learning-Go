package main

import "fmt"

func main() {
}

type e interface {}

func mult2 (f e) e {
    switch type(f) {
    case int:
        return int(f) * 2
    }
    return f
}
