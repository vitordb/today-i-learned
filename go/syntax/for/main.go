package main

import "fmt"

func main() {

	var animals []string

	animals = append(animals, "cat", "dog", "shark")

	for i := 0; i < len(animals); i++ {

		fmt.Println(animals[i])

	}
}
