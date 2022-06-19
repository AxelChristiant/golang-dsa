package main

import "fmt"

func partition(arr *[]int, low int, high int) int {
	pivot := (*arr)[high]
	i := low - 1
	for j := low; j < high; j++ {
		if (*arr)[j] <= pivot {
			i = i + 1
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}
	(*arr)[i+1], (*arr)[high] = (*arr)[high], (*arr)[i+1]

	return i + 1

}

func quickSort(arr *[]int, low int, high int) {
	if low < high {
		pi := partition(&(*arr), low, high)
		quickSort(&(*arr), low, pi-1)
		quickSort(&(*arr), pi+1, high)

	}
}

func main() {
	x := []int{6, 7, 3, 2, 1, 4}
	quickSort(&x, 0, len(x)-1)
	fmt.Println(x)

}
