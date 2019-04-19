//To remove an element from the middle of a slice, preserving the order of the remaining ele- ments, use copy to slide the
//higher-numbered elements down by one to fill the gap:

package main

import "fmt"

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 8 9]"
}

func removeWithoutPreservingOrder(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
