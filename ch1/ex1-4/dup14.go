// dup14 prints the count of text of lines that appear more
// than once in the provided files and te names of the files
// in which the duplicate lines appear

// go run dup14.go sample1.txt sample2.txt sample3.txt
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	inFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("please provide file names")
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup14: %v\n", err)
				continue
			}
			countLines(f, counts, inFiles, file)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, inFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, inFiles map[string][]string, file string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		str := input.Text()
		counts[str]++
		if !contains(inFiles[str], file) {
			inFiles[str] = append(inFiles[str], file)
		}
	}
	// ignoring input.Err()
}

func contains(list []string, str string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}
