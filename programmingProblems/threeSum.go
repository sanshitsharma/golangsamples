package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
)

func threeSumBF(arr []int) ([][]int, error) {
	if arr == nil || len(arr) == 0 {
		return nil, errors.New("invalid params")
	}

	res := make([][]int, 0)
	for i := 0; i < len(arr) - 2; i++ {
		for j := i + 1; j < len(arr) - 1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i] + arr[j] + arr[k] == 0 {
					res = append(res, []int{arr[i], arr[j], arr[k]})
				}
			}
		}
	}

	return res, nil
}

func threeSumOptimized(arr []int) ([][]int, error) {
	if arr == nil || len(arr) == 0 {
		return nil, errors.New("invalid params")
	}

	res := make([][]int, 0)
	sort.Ints(arr)

	for i := 0; i < len(arr) - 2; i++ {
		j := i + 1
		k := len(arr) - 1

		trgt := 0 - arr[i]

		for j < k {
			sum := arr[j] + arr[k]
			if sum == trgt {
				res = append(res, []int{arr[i], arr[j], arr[k]})
				j++
				k--
			} else if sum < trgt {
				j++
			} else {
				k--
			}
		}
	}

	return res, nil
}

func main() {
	fmt.Println("Finding 3 sum.")

	arr := []int{4, 3, -1, 2, -2, 10, -7, -3, 0}

	fmt.Println("arr:", arr)

	if triplets, err := threeSumBF(arr); err != nil {
		fmt.Printf("failed to find 3 sum. err: %v", err)
		os.Exit(1)
	} else {
		fmt.Printf("Result Brute Force: %v\n", triplets)
	}

	if triplets, err := threeSumOptimized(arr); err != nil {
		fmt.Printf("failed to find 3 sum. err: %v", err)
		os.Exit(1)
	} else {
		fmt.Printf("Result Optimized: %v\n", triplets)
	}
}
