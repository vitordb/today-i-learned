package main

import (
	"fmt"
)

func main() {

	//basic array usage and declaration, fixed length, you can't change it
	var ages = [3]int{20, 25, 30}

	//short hand
	names := [4]string{"Vitor", "Josh", "Mario", "John"}
	names[1] = "Megan"

	//slice
	animals := []string{"cat", "dog", "mouse", "lion"}
	animals[3] = "tiger"
	animals = append(animals, "horse")

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))
	fmt.Println(animals, len(animals))

	newArray := []int{1, 2, 3, 4, 5}

	fmt.Println(newArray[2])

	fmt.Println(len(newArray))

	//creating an empty slice
	names2 := []string{}

	fmt.Println(names2)

	//using the make builtin function
	s := make([]string, 3)
	fmt.Println("emp", s, "len", len(s), "cap", cap(s))

	//copying a slice into another slice
	c := make([]string, 2)
	copy(c, s)
	fmt.Println("copied slice", c)

	//removing elements
	newArray = append(newArray[:1], newArray[2:]...)

	fmt.Println(newArray)

}
