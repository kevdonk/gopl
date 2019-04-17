/*
Exercise 2.5: The expression x&(x-1) clears the rightmost non-zero bit of x.
Write a version of PopCount that counts bits by using this fact, and assess its performance.
*/

package main

import "fmt"

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	total := 0
	for x != 0 {
		x = x & (x - 1)
		total++
	}
	return total
}

func main() {
	fmt.Println(PopCount(0x1234567890ABCDEF))
}
