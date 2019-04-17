/*
Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression. Compare the per- formance of the two versions.
(Section 11.4 shows how to compare the performance of differ- ent implementations systematically.)
*/

package main

import (
	"fmt"

	"github.com/kevdonk/gopl/ch2/ex2.3/popcount"
)

func main() {
	fmt.Println(popcount.PopCount(4))
}
