// Exercise 4.9: Write a program wordfreq to report the frequency of each word in an input text file.
// Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into words instead of lines.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
			continue
		}
		defer f.Close()
		countWords(f)
	}
}

func countWords(f *os.File) {
	counts := map[string]int{}
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Print("word\tcount\n")
	for s, i := range counts {
		fmt.Printf("%s\t%d\n", s, i)
	}
}
