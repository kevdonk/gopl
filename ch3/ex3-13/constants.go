// Exercise 3.13: Write const declarations for KB, MB, up through YB as compactly as you can.
package main

import "fmt"

const (
	B  = float64(1)
	KB = 1000 * B
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main() {
	fmt.Printf(" B: %v\nKB: %v\nMB: %v\nGB: %v\nTB: %v\nPB: %v\nEB: %v\nZB: %v\nYB: %v\n",
		B, KB, MB, GB, TB, PB, EB, ZB, YB)
}
