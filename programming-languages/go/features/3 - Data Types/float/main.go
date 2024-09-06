package main

import "fmt"

/*
The float data types are used to store positive and negative numbers with a decimal point, like 35.3, -2.34, or 3597.34987.
The float data type has two keywords: float32 and float64.
The float32 data type can store numbers with a decimal point that have up to 7 digits of precision,
while the float64 data type can store numbers with a decimal point that have up to 15 digits of precision.
*/
func main() {

	var x float64 = 1.7e+308
	fmt.Printf("Type: %T, value: %v", x, x)

}
