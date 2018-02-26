package main

import (
	"fmt"
	"errors"
)

func partition(a []int, l, h int) int {
	pivot := a[h]

	for l < h {
		for a[l] < pivot {
			l++
		}

		for a[h] > pivot {
			h--
		}

		temp := a[l]
		a[l] = a[h]
		a[h] = temp
	}

	return l
}

func qSelect(a []int, l, h, k int) int {
	for l < h {
		pivotIndx := partition(a, l, h)
		if pivotIndx == k - 1 {
			return a[pivotIndx]
		} else if (pivotIndx > k - 1) {
			return qSelect(a, l, pivotIndx - 1, k)
		} else {
			return qSelect(a, pivotIndx + 1, h, k)
		}
	}

	return -1
}

func QuickSelect(a []int, k int) (int, error) {
	if k > len(a) {
		return -1, errors.New("invalid kth index")
	}

	l := 0
	h := len(a) - 1

	return qSelect(a, l, h, k), nil
}

func main() {
	fmt.Println("Find the kth smallest element in an unordered list")
	a := []int{2, 7, 1, 3, 6, 4}
	k := 1

	if kthSmallest, err := QuickSelect(a, k); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%dth smallest element in %v = '%d'\n", k, a, kthSmallest)
	}
}