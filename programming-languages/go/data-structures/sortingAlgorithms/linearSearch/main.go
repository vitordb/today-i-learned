package main

import "fmt"

func main() {

	slice := []int{0, 2, 6, 7, 10, 12}

	linearSearch(slice, 7)

}

func linearSearch(slice []int, target int) int {

	for _, v := range slice {

		if target == v {
			fmt.Println("value is:", v)
			return v
		}

	}

	return -1

}
