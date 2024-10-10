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

	sliceCloned := slices.Clone(slice3)
	fmt.Println("slice cloned", sliceCloned)

	//return index
	fmt.Println(slices.Index(slice3, 3))

	//insert

	sliceCloned = slices.Insert(sliceCloned, 2, 1, 2, 3, 4)
	fmt.Println(sliceCloned)

	//append with slice package
	a := []int{1, 2, 3}
	seq := slices.Values([]int{4, 5})
	c := slices.AppendSeq(a, seq)
	fmt.Println(c)

}
