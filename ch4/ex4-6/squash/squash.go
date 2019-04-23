// Exercise 4.6: Write an in-place function that squashes each run of adjacent Unicode spaces (see unicode.IsSpace)
// in a UTF-8-encoded []byte slice into a single ASCII space.

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squash(s []byte) []byte {
	read := 0
	write := 0
	prevIsSpace := false
	for read < len(s) {
		r, size := utf8.DecodeRune(s[read:])
		if !unicode.IsSpace(r) {
			utf8.EncodeRune(s[write:], r)
			write += size
			prevIsSpace = false
		} else {
			if !prevIsSpace {
				size = utf8.EncodeRune(s[write:], rune(' '))
				write += size
			}
			prevIsSpace = true
		}
		read++
	}
	return s[:write]
}

func main() {
	b := []byte{'t', 'o', 'o', '\t', '\n', '\v', '\f', '\r', ' ', 'm', 'u', 'c', 'h', ' ', ' ', 'n', 'o', 'w'}
	b = squash(b)
	fmt.Println(string(b))
}
