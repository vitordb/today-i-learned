package main

import (
	"fmt"
	"slices"
)

func main() {

	//slices packages
	slice1 := []string{"test1", "test2", "test3"}
	slice2 := []string{}

	slice2 = append(slice2, "test1", "test2", "test3")
	//comparing slices
	if slices.Equal(slice1, slice2) {
		fmt.Println("slice1 == slice2")
	}

	//cloning slice
	slice3 := []int{1, 2, 3}

	fmt.Println(slice3Copy)

}
