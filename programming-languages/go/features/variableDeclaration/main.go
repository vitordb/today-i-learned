package main

import "fmt"

// two ways to declare variable

// with the var keyword
var example string = "Dog"

// need to use the variable
func main() {

	// with the := sign
	animal := "Cat"

	fmt.Println("animal: ", animal)

	// value assignment after declaration
	var student1 string
	student1 = "josh"
	fmt.Println(student1)

}
