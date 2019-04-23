// Exercise 4.4: Write a version of rotate that operates in a single pass.
package main

import "fmt"

// rotate rotates slice 's' left by 'n' elements
func rotate(s []int, n int) {
	if n > len(s) {
		rotate(s, n-len(s))
	} else {
		tmp := make([]int, n)
		copy(tmp, s[:n])
		copy(s, s[n:])
		copy(s[len(s)-n:], tmp)
	}
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 8)
	fmt.Printf("%v\n", s) // 2, 3, 4, 5, 0, 1
}
