package main

import "fmt"

func insertionSort(arr *[]int) {
	for j := 1; j < len(*arr); j++ {
		key := (*arr)[j]

		i := j - 1
		for {
			if i > -1 && (*arr)[i] > key {
				(*arr)[i+1] = (*arr)[i]
				i--
			} else {
				break
			}
		}
		(*arr)[i+1] = key
	}

}

func main() {
	tes := []int{4, 2, 3, 1}
	insertionSort(&tes)
	fmt.Println(tes)
}
