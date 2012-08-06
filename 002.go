package main

import (
    "fmt"
)

func main() {
    a, b, s := 1, 2, 2
    for b < 4000000 {
        b = b + a
        a = b - a
        if b % 2 == 0 {
            s += b
        }
    }
    fmt.Printf("%d", s)
 
}
