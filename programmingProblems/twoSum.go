package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
)

// Given a array of integers find all the pairs whose
// sum is equal to 0

func twoSumWithMap(arr []int) ([][]int, error) {
	if arr == nil || len(arr) < 2 {
		return nil, errors.New("invalid input")
	}

	res := make([][]int, 0)
	sumMap := make(map[int]bool)

	for _, i := range arr {
		if sumMap[0-i] {
			pair := []int{i, 0-i}
			res = append(res, pair)
		} else {
			sumMap[i] = true
		}
	}

	return res, nil
}

func twoSum(arr []int) ([][]int, error) {
	if arr == nil || len(arr) < 2 {
		return nil, errors.New("invalid input")
	}

	sort.Ints(arr)
	//fmt.Println("sorted input:", arr)

	res := make([][]int, 0)
	i := 0
	j := len(arr) - 1
	for i < j {
		sum := arr[i] + arr[j]
		if sum == 0 {
			res = append(res, []int{arr[i], arr[j]})
			i++
			j--
		} else if sum < 0 {
			i++
		} else {
			j--
		}
	}

	return res, nil
}

func main() {
	fmt.Println("Finding 2 Sum")
	arr := []int{1,-1,0,2,-2}

	fmt.Printf("Input Arr: %v\n", arr)

	if pairs, err := twoSumWithMap(arr); err != nil {
		fmt.Println("failed to find 2Sum pairs. err:", err)
		os.Exit(1)
	} else {
		fmt.Printf("Result with Map: %v\n", pairs)
	}

	if pairs, err := twoSum(arr); err != nil {
		fmt.Println("failed to find 2Sum pairs. err:", err)
		os.Exit(1)
	} else {
		fmt.Printf("Result without Map: %v\n", pairs)
	}
}