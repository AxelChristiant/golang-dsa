package main

import (
	"fmt"
)

func binarySearch(arr *[]int, q int) int {
	high := len(*arr) - 1
	low := 0
	for low < high {
		mid := (high - low) / 2
		if (*arr)[mid] < q {
			low = mid - 1
		} else if (*arr)[mid] > q {
			high = mid + 1
		} else {
			return mid
		}
	}
	return -1

}

func main() {
	tes := []int{1, 2, 3, 4, 5}
	fmt.Println(binarySearch(&tes, 3))

}
