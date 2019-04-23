// Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.

package main

import "fmt"

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func dedupe(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			s = remove(s, i)
		}
	}
	return s
}

func main() {
	s := []string{"hey", "hey", "you", "you", "get", "off", "of", "my", "cloud", "cloud"}
	s = dedupe(s)
	fmt.Println(s)
}
