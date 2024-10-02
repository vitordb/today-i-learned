package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1

}

func main() {

	abs := make([]int, 20)

	abs = append(abs, 1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)

	bs := BinarySearch(abs, 13)

	fmt.Println(bs)

}
