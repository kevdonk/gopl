/*
Exercise 2.4: Write a version of PopCount that counts bits by shifting its argument through 64 bit positions,
testing the rightmost bit each time. Compare its performance to the table- lookup version.
*/
package main

import "fmt"

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	total := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			total++
		}
	}
	return total
}

func main() {
	fmt.Println(PopCount(0x1234567890ABCDEF))
}
