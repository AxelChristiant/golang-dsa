package main

import (
	"fmt"
	"math"
)

type fibonacciTree struct {
	value int
	child []fibonacciTree
	order int
}

type fibonacciHeap struct {
	trees []fibonacciTree
	least fibonacciTree
	count int
}

func NewFibonacciHeap() fibonacciHeap {
	return fibonacciHeap{least: fibonacciTree{value: int(math.Inf(-1))}}
}

func (f *fibonacciTree) addAtTheEnd(anotherTree fibonacciTree) {
	f.child = append(f.child, anotherTree)
	f.order += 1
}

func (fh *fibonacciHeap) insertNode(value int) {
	new_tree := fibonacciTree{value: value}
	fh.trees = append(fh.trees, new_tree)
	if fh.least.value == int(math.Inf(-1)) || value < fh.least.value {
		fh.least = new_tree
	}
}

func (fh *fibonacciHeap) getMin() int {
	return fh.least.value
}
func (fh *fibonacciHeap) extractMin() int {
	smallest := fh.least
	if smallest.value != int(math.Inf(-1)) {
		for _, child := range smallest.child {
			fh.trees = append(fh.trees, child)
		}
		if len(fh.trees) == 0 {
			fh.least.value = int(math.Inf(-1))
		} else {
			fh.least = fh.trees[0]
			return 0

		}
	}
	return smallest.value
}

// func (fh *fibonacciHeap) consolidate() {
// 	_,degree := math.Frexp(float64(fh.count)) + 1
// 	aux
// 	order :=

// }
func main() {
	// tes := fibonacciTree{value: 10}
	// tes2 := fibonacciTree{value: 20}
	// tes.addAtTheEnd(tes2)
	tes := NewFibonacciHeap()
	fmt.Println(tes)

}
