// Exercise 3.10: Write a non-recursive version of comma, using bytes.Buffer instead of string concatenation.
// comma inserts commas in a non-negative decimal integer string.
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	for i, r := range s {
		buf.WriteRune(r) // Use WriteRune for arbitrary UTF-8
		remaining := len(s) - (i + 1)
		if remaining > 0 && remaining%3 == 0 {
			buf.WriteByte(',') // WriteByte okay for ascii
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("121212121"))
	fmt.Println(comma("121"))
	fmt.Println(comma("2121"))
}
