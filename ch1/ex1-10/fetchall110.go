// fetchall fetches URLs in parallel and reports their times and sizes
/*
Exercise 1.10: Find a web site that produces a large amount of data. Investigate caching by running fetchall twice
in succession to see whether the reported time changes much. Do you get the same content each time? Modify fetchall
to print its output to a file so it can be examined.
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	f, err := os.Create("dump.html")
	if err != nil {
		fmt.Printf("fetchall110: error creating file: %v\n", err)
	}
	defer f.Close()
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading: %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
