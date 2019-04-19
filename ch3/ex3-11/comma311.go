// comma inserts commas in a non-negative decimal integer string.
package main

import (
	"fmt"
	"strings"
)

func comma(s string) string {
	offset := 0
	if s[0] == '+' || s[0] == '-' {
		offset = 1
	}
	decimal := ""
	d := strings.LastIndex(s, ".")
	if d > -1 {
		decimal = s[d:]
		s = s[:d]
	}
	n := len(s)
	if n <= 3+offset {
		return s + decimal
	}
	return comma(s[:n-3]) + "," + s[n-3:] + decimal
}

func main() {
	fmt.Println(comma("121212121"))
	fmt.Println(comma("-121"))
	fmt.Println(comma("2121"))
	fmt.Println(comma("+121212121.565675"))
	fmt.Println(comma("-121.567"))
	fmt.Println(comma("2121.0"))
}
