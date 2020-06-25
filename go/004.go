/*
Problem 4

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
)

func is_palindrome(n int) bool {
	a := 0
    for n > 0 {
        a = a * 10 + n % 10
        n /= 10
        if (a == n) {
			return true
		}
	}
    return false
}

func main() {
	n := 0
	for a := 100; a < 1000; a++ {
		for b:= 100; b < 1000; b++ {
			if is_palindrome(a * b) && a * b > n {
				n = a * b
			}
		}
	}
	fmt.Println(n)
}