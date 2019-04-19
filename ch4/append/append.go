package append

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	copy(z[len(x):], y)
	return z
}

// func main() {
// 	var x []int
// 	x = appendInt(x, 1)
// 	x = appendInt(x, 2, 3)
// 	x = appendInt(x, 4, 5, 6)
// 	x = appendInt(x, x...) // append the slice x
// 	fmt.Println(x)         // "[1 2 3 4 5 6 1 2 3 4 5 6]"
// }
