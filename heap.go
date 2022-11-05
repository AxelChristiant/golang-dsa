package main

import (
	"fmt"
	"math"
)

func heapify(arr *[]int, n int, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	if l < n && (*arr)[i] < (*arr)[l] {
		largest = l
	}
	if r < n && (*arr)[largest] < (*arr)[r] {
		largest = r
	}
	if largest != i {
		(*arr)[i], (*arr)[largest] = (*arr)[largest], (*arr)[i]
		heapify(arr, n, largest)
	}
}

func insert(arr *[]int, newNum int) {
	size := len((*arr))
	if size == 0 {
		(*arr) = append((*arr), newNum)
	} else {
		(*arr) = append((*arr), newNum)
		for i := (math.Floor(float64(size / 2))); i > -1; i-- {
			heapify(arr, size, int(i))
		}
	}

}

func deleteNode(arr *[]int, num int) {
	size := len((*arr))
	for i := 0; i < size; i++ {
		if num == (*arr)[i] {
			(*arr)[i], (*arr)[size-1] = (*arr)[size-1], (*arr)[i]
			break
		}
	}
	(*arr) = RemoveIndex((*arr), size-1)
	size = len((*arr))
	for i := (math.Floor(float64(size / 2))); i > -1; i-- {
		heapify(arr, size, int(i))
	}
}

func RemoveIndex(arr []int, indexToRemove int) []int {
	return append(arr[:indexToRemove], arr[indexToRemove+1:]...)
}

func main() {
	tes := make([]int, 0)
	insert(&tes, 3)
	insert(&tes, 4)
	insert(&tes, 9)
	insert(&tes, 5)
	insert(&tes, 2)
	fmt.Println(tes)
	deleteNode(&tes, 4)
	fmt.Println(tes)
}
