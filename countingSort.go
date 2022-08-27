package main

import "fmt"

func countingSort(arr *[]int) {
	size := len(*(arr))
	output := make([]int, size)
	count := make([]int, 10)

	for i := 0; i < size; i++ {
		count[(*arr)[i]] += 1
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	i := size - 1
	for i >= 0 {
		output[count[(*arr)[i]]-1] = (*arr)[i]
		count[(*arr)[i]] -= 1
		i -= 1
	}
	for i := 0; i < size; i++ {
		(*arr)[i] = output[i]
	}
}

func main() {
	x := []int{6, 7, 3, 2, 1, 4}
	countingSort(&x)
	fmt.Println(x)

}
