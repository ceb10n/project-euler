/*
Problem 3.

The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

package main

import (
	"fmt"
)

func main() {
	var i int64 = 2
	var n int64 = 600851475143
	for n != 1 {
		if n % i == 0 {
			n /= i
		} else {
			i += 1
		}
	}
	
	fmt.Printf("%d", i)
}