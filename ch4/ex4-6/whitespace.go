// Exercise 4.6: Write an in-place function that squashes each run of adjacent Unicode spaces (see unicode.IsSpace)
// in a UTF-8-encoded []byte slice into a single ASCII space.

package main

import (
	"fmt"
	"unicode"
)

// remove removes ith element from a slice
func remove(slice []byte, i int) []byte {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// insert adds byte b to slice at index i
func insert(slice []byte, i int, b byte) []byte {
	// our example will always have capacity, because we only add after a removal

	// not really in-place...
	tmp := make([]byte, len(slice)-i)
	copy(tmp, slice[i:])
	slice = append(slice[:i], b)
	slice = append(slice, tmp...)
	return slice
}

func trim(s []byte) []byte {
	for i := 0; i < len(s); i++ {
		spaces := 0
		for j := 0; j < len(s); j++ {
			if unicode.IsSpace(rune(s[i])) {
				spaces++
				s = remove(s, i)
			}
		}
		if spaces >= 1 {
			s = insert(s, i, ' ')
		}
	}
	return s
}

func main() {
	b := []byte{'t', 'o', 'o', '\t', '\n', '\v', '\f', '\r', ' ', 'm', 'u', 'c', 'h', 0x85, 0xA0, 'n', 'o', 'w'}
	b = trim(b)
	fmt.Println(string(b))
}
