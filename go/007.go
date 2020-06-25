package main

import (
	"fmt"
)



func main() {
	is_prime := func(n int) bool {
					if n == 2 || n == 3 { return true }
					for i := 2; i < n / 2; i++ {
						if n % i == 0 { return false }
					}
					return true
				}
	c, p := 0, 0
	for i := 2; c <= 10001; i++ {
		if is_prime(i) { 
			c += 1
			p = i
		}
	}
	fmt.Println(p);
}