package main

import "fmt"

func main() {

	/* two types of integers:
	Signed integers: can store both positive and negative values
	Unsigned integers: can only store non-negative values
	*/

	//Default type for integer is int

	// Signed integers

	/*
		go has five keywords/types of signed integers:
		int     depends on platform, 31 bits, or 64 bits   range of-2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems
		int8	8 bits/1 byte	-128 to 127
		int16	16 bits/2 byte	-32768 to 32767
		int32	32 bits/4 byte	-2147483648 to 2147483647
		int64	64 bits/8 byte	-9223372036854775808 to 9223372036854775807int8 - 8 bits/1 byte	-128 to 127

	*/

	var x int = 500
	y := -4500

	fmt.Printf("Type: %T, value: %v\n", x, x)

	fmt.Printf("Type: %T, value: %v\n", y, y)

}
