package main

import (
	"fmt"

	qs "github.com/sanshitsharma/golangsamples/sorting/quick_sort"
)

func quickSortSample() {
	fmt.Println("Example of quicksort")

	//a := []int{10, 9, 8, 7, 5, 1}
	a := []int{1, 5, 7, 8, 9, 10}
	qs.Sort(a)

	fmt.Println(a)
}

func main() {
	quickSortSample()
}
