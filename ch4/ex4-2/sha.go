//Exercise 4.2: Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag to print the SHA384 or SHA512 hash instead.

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var b = flag.Int("b", 256, "hash rate (256, 384, or 512 bit)")

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		switch *b {
		case 256:
			fmt.Printf("SHA%v (%v): %x\n", *b, arg, sha256.Sum256([]byte(arg)))
		case 512:
			fmt.Printf("SHA%v (%v): %x\n", *b, arg, sha512.Sum512([]byte(arg)))
		case 384:
			fmt.Printf("SHA%v (%v): %x\n", *b, arg, sha512.Sum384([]byte(arg)))
		default:
			fmt.Println("Invalid value for flag b: 256, 384 or 512, please")
		}
	}
}
