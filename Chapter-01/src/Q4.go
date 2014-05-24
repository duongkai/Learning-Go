package main

import (
    "fmt"
    "unicode/utf8"
)

func q4_1() {
    str := ""
    for gi := 0; gi < 100; gi++ {
        str += "A"
        fmt.Println (str)
    }
    //return 0
}

func q4_2() {
    str := "asSASA ddd dsjkdsjs dk tôi là"
    //str := "日本語"
    fmt.Println ("length of the string = ", len (str))
    /*
    nLen := 0
    for gc, gd := range str {
        fmt.Println (gd)
        a, byteLen := utf8.DecodeRuneInString (str[gc:gc+1])
        fmt.Println (a)
        nLen += byteLen
    }
    fmt.Println ("number of bytes = ", nLen)
    */
    fmt.Println ("number of bytes = ", utf8.RuneCountInString (str))
}

func q4_2_2() {
    str := "tôi l phạm tùng dương"
    fmt.Printf ("String %s \n Length: %d, Runes: %d\n", str, len (str),
        utf8.RuneCountInString (str))
}

func ReverseStr (s string) string {
    out := make ([]rune, utf8.RuneCountInString(s))
    lin := len (out)
    for _, c := range s {
        lin--
        out[lin] = c
    }
    return string (out)
}

func main() {
    //q4_1()
    //q4_2()
    //fmt.Println (ReverseStr ("tôi là phạm tùng dương"))
    q4_2_2()
}
