package main

import (
    "fmt"
    "math/big"
)

func main() {
    fmt.Println (factorial (10000))
}

func factorial (n int64) *big.Int {
    res := big.NewInt (1)
    if n == 0 {
        return big.NewInt(1)
    } else {
        /*for i = 1; i <= n; i++ {
            tmp := big.NewInt (i)
            res.Mul (res, tmp)
        }*/
        res.MulRange (1, n) 
        return res
    }
}   
