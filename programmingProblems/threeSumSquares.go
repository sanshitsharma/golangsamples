package main

import (
	"fmt"
	"errors"
	"os"
)

func threeSumSquares(arr []int) ([][]int, error){
	if arr == nil || len(arr) < 3 {
		return nil, errors.New("invalid params")
	}

	sumMap := make(map[int][][]int, 0)

	for i := 0; i < len(arr) - 1; i++ {
		for j := i + 1; j < len(arr); j++ {
			sum := arr[i]*arr[i] + arr[j]*arr[j]
			if sumMap[sum] == nil {
				sumMap[sum] = [][]int{{arr[i], arr[j]}}
			} else {
				sumMap[sum] = append(sumMap[sum], []int{arr[i], arr[j]})
			}
		}
	}

	res := make([][]int, 0)
	for i := range arr {
		trgt := arr[i]*arr[i]
		if sumMap[trgt] != nil {
			for _, items := range sumMap[trgt] {
				items = append(items, arr[i])
				res = append(res, items)
			}
		}
	}

	return res, nil
}

func main() {
	fmt.Println("Given a list of integers, find all triplets which satisfy a^2 + b^2 = c^2")
	arr := []int{3, -63, -13, -5, -3, 4, 7, 12, 16, 65, 1, 2}

	res, err := threeSumSquares(arr)
	if err != nil {
		fmt.Printf("failed to find square triplets.. err: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("pythagorean triplets:", res)
}