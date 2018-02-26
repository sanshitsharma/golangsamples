package main

import (
	"fmt"

	qs "github.com/sanshitsharma/golangsamples/sorting/quick_sort"
)

func quickSortSample() {
	fmt.Println("Example of quicksort")

	a := []int{2, 7, 1, 3, 6, 4}
	qs.Sort(a)

	fmt.Println(a)
}

func main() {
	quickSortSample()
}
