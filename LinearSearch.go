package main

import (
	"fmt"
)

func binarySearch(arr *[]int, q int) int {
	for i := 0; i < len((*arr)); i++ {
		if (*arr)[i] == q {
			return i
		}

	}
	return -1

}

func main() {
	tes := []int{1, 2, 3, 4, 5}
	fmt.Println(binarySearch(&tes, 3))

}
