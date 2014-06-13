package main

import "fmt"

func main() {
    var p *int
    i := 10
    p = &i
    fmt.Println ("*p = ", *p)
    // Change value
    fmt.Println ("Changing value...")
    i++
    fmt.Println ("*p = ", *p)
    // Change value of p
    fmt.Println ("Change value of p")
    *p = 15
    fmt.Println ("i = ", i)
}
