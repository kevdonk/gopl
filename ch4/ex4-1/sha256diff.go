// Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256 hashes.
// (SeePopCountfromSection2.6.2.)

package main

import (
	"crypto/sha256"
	"fmt"
)

func shaDiff(a, b [32]uint8) int {
	count := 0
	for i := 0; i < 32; i++ {
		count += byteDiff(a[i], b[i])
	}
	return count
}

// byteDiff returns the number of bits that are different in two bytes
func byteDiff(a, b uint8) int {
	count := 0
	for i := uint(0); i < 8; i++ {
		if a&(1<<i) != b&(1<<i) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(shaDiff(sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X"))))
	fmt.Println(byteDiff(byte('a'), byte('A')))
	fmt.Println(byteDiff(byte('a'), byte('a')))
	fmt.Println(byteDiff(byte('b'), byte('A')))
}
