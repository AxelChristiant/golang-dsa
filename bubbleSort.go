package main

import (
	"fmt"
)

func bubbleSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}

}

func main() {
	tes := []int{4, 2, 3, 1}
	bubbleSort(&tes)
	fmt.Println(tes)
}
