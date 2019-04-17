// Fetch prints the content found at a URL
/*
Exercise1.7:Thefunctioncallio.Copy(dst,src)readsfromsrcandwritestodst. Useit instead of ioutil.ReadAll to copy the response
body to os.Stdout without requiring a buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
