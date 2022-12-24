package main

import (
	"fmt"
	"math"
	"reflect"
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
	fh.count += 1
}

func (fh *fibonacciHeap) deleteSmallest() {
	i := 0
	for _, tree := range fh.trees {
		if !reflect.DeepEqual(tree, fh.least) {
			fh.trees[i] = tree
			i += 1
		}
	}
	fh.trees = fh.trees[:i]

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
		fh.deleteSmallest()
		if len(fh.trees) == 0 {
			fh.least.value = int(math.Inf(-1))
		} else {
			fh.least = fh.trees[0]
			fh.consolidate()
		}
		fh.count = fh.count - 1
	}
	return smallest.value
}

func (fh *fibonacciHeap) consolidate() {
	_, degree := math.Frexp(float64(fh.count))
	degree = degree
	aux := make([]fibonacciTree, degree)
	for len(fh.trees) > 0 {
		x := fh.trees[0]
		order := x.order
		fh.trees = append(fh.trees[:0], fh.trees[1:]...)
		for aux[order].value > 0 {
			y := aux[order]
			if x.value > y.value {
				x, y = y, x
			}
			x.addAtTheEnd(y)
			aux[order] = fibonacciTree{}
			order = order + 1
		}
		aux[order] = x
	}
	fh.least = fibonacciTree{}
	for _, k := range aux {
		if k.value != 0 {
			fh.trees = append(fh.trees, k)
			if (fh.least.value == 0) || (k.value < fh.least.value) {
				fh.least = k
			}
		}

	}

}
func main() {
	fibonacciheap := NewFibonacciHeap()
	fibonacciheap.insertNode(7)
	fibonacciheap.insertNode(3)
	fibonacciheap.insertNode(17)
	fibonacciheap.insertNode(24)
	fibonacciheap.extractMin()
	fmt.Println(fibonacciheap.getMin())

}
