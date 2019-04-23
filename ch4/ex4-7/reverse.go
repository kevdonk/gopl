// Exercise 4.7: Modify reverse to reverse the characters of a []byte slice that represents a UTF-8-encoded string, in place.
// Can you do it without allocating new memory?

package main

import (
	"fmt"
	"unicode/utf8"
)

// reverse reverses a slice of ints in place.
func reverse(s []byte) {
	var first, last rune
	var fSize, lSize int
	for i, j := 0, len(s); i < j-1; i, j = i+lSize, j-fSize {
		first, fSize = utf8.DecodeRune(s[i:])
		last, lSize = utf8.DecodeLastRune(s[:j])

		//if lSize > fSize, need to move all of the bytes of the last rune to the front
		if lSize > fSize {
			copy(s[i+lSize:], s[i+fSize:j-lSize])
		}

		copy(s[i:], []byte(string(last)))
		copy(s[j-fSize:], []byte(string(first)))
	}
}

func main() {
	s := []byte{65, 66, 67, 226, 130, 172} // ABC€
	reverse(s)
	fmt.Println(string(s))
}
