/*
Go also provides the following bitwise binary operators, the first four of which treat their op-
erands as bit patterns with no concept of arithmetic carry or sign:
& bitwise AND
| bitwise OR
^ bitwise XOR
&^ bit clear (AND NOT) << left shift
>> right shift
The operator ^ is bitwise exclusive OR (XOR) when used as a binary operator, but when used as a unary prefix operator it is bitwise negation or complement; that is, it returns a value with each bit in its operand inverted. The &^ operator is bit clear (AND NOT): in the expression z = x &^ y, each bit of z is 0 if the corresponding bit of y is 1; otherwise it equals the cor- responding bit of x.

The code below shows how bitwise operations can be used to interpret a uint8 value as a compact and efficient set of 8 independent bits. It uses Printf’s %b verb to print a number’s binary digits; 08 modifies %b (an adverb!) to pad the result with zeros to exactly 8 digits.
*/
package bitwise

import "fmt"

var x uint8 = 1<<1 | 1<<5
var y uint8 = 1<<1 | 1<<2

func main() {
	fmt.Printf("%08b\n", x)    // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y)    // "00000110", the set {1, 2}
	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}
	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}

/*
(Section 6.5 shows an implementation of integer sets that can be much bigger than a byte.)
In the shift operations x<<n and x>>n, the n operand determines the number of bit positions to shift and must be unsigned; the x operand may be unsigned or signed. Arithmetically, a left shift x<<n is equivalent to multiplication by 2n and a right shift x>>n is equivalent to the floor of division by 2n.
Left shifts fill the vacated bits with zeros, as do right shifts of unsigned numbers, but right shifts of signed numbers fill the vacated bits with copies of the sign bit. For this reason, it is important to use unsigned arithmetic when you’re treating an integer as a bit pattern.
*/

// see if ith bit is set
// (x & 1<<i == 0)

// set ith bit
// x |= 1<<i
// or
// x &= ^(1<<i)

// clear ith bit
// x &= 1<<i

// toggle ith bit
// x ^= 1<<i

// hex / octal
// o := 0666
// fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
// x := int64(0xdeadbeef)
// fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
// Output:
// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
