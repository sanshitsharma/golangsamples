package main

import (
	"errors"
	"fmt"

	"github.com/sanshitsharma/golangsamples/ds/heap"
)

func findMedian(list []int, k int) (int, error) {
	lo, _ := heap.NewHeap(heap.Max)
	hi, _ := heap.NewHeap(heap.Min)

	// Process elements till index k
	for i := 0; i <= k; i++ {
		//fmt.Println("processing:", list[i])
		elem := list[i]

		// First insert to lo heap
		lo.Insert(elem)

		// if size diff between lo and hi > 1, then offer the top element in lo to hi
		if lo.Size()-hi.Size() > 1 {
			loTop, _ := lo.Pop()
			hi.Insert(loTop)
		}

		// if size of hi has become greater than the size of lo after the last offering
		// then insert the top element in hi to lo
		if hi.Size() > lo.Size() {
			hiTop, _ := hi.Pop()
			lo.Insert(hiTop)
		}
	}

	// If the number of elements processed is odd, then return the top of lo
	// else return the mean of top of lo & hi
	if !lo.IsEmpty() {
		if (lo.Size()+hi.Size())%2 != 0 {
			return lo.Peek()
		} else {
			loTop, _ := lo.Peek()
			hiTop, _ := hi.Peek()
			return (loTop + hiTop) / 2, nil
		}
	}

	return -1, errors.New(`failed to find median`)
}

func main() {
	fmt.Println(`Given a streaming list of integers find the median at a given point`)
	list := []int{5, 15, 1, 3, 2, 8, 7, 9, 10, 6, 11, 4}
	k := len(list) - 1

	if median, err := findMedian(list, k); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("median for %v at index %v = '%v'\n", list[0:k+1], k, median)
	}
}
