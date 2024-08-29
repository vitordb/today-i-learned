package main

import "fmt"

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

	//removing elements
	newArray = append(newArray[:3], newArray[4:]...)

	fmt.Println(newArray)
}
