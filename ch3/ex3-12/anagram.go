// Exercise 3.12: Write a function that reports whether two strings are anagrams of each other, that is,
// they contain the same letters in a different order.

package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("abba baba: %v\n", areAnagrams("abba", "baba"))
	fmt.Printf("abba nono: %v\n", areAnagrams("abba", "nono"))
	fmt.Printf("race car car race: %v\n", areAnagrams2("race car", "car race"))
	fmt.Printf("racecar carrace: %v\n", areAnagrams2("racecar", "car race"))
}

func areAnagrams(a, b string) bool {
	aBytes := []byte(a)
	bBytes := []byte(b)
	sort.Slice(aBytes, func(i, j int) bool { return aBytes[i] < aBytes[j] })
	sort.Slice(bBytes, func(i, j int) bool { return bBytes[i] < bBytes[j] })
	if bytes.Compare(aBytes, bBytes) == 0 {
		return true
	}

	return false
}

func areAnagrams2(a, b string) bool {
	countA := 0
	for _, c := range a {
		countA += int(c)
	}
	countB := 0
	for _, c := range b {
		countB += int(c)
	}
	if countA == countB {
		return true
	}
	return false
}
