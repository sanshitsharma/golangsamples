package main

import (
	"fmt"
	"time"
)

/*
Catalan numbers are a sequence of natural numbers that occurs in many interesting counting problems like following.

1) Count the number of expressions containing n pairs of parentheses which are correctly matched. For n = 3, possible
expressions are ((())), ()(()), ()()(), (())(), (()()).

2) Count the number of possible Binary Search Trees with n keys (See this)

3) Count the number of full binary trees (A rooted binary tree is full if every vertex has either two children or
no children) with n+1 leaves.
*/

func nthCatalanNumber(n int) uint64 {
	 if n <= 1 {
	 	return 1
	 }

	 var res uint64
	 for i := 0; i < n; i++ {
	 	res += nthCatalanNumber(i) * nthCatalanNumber(n-i-1)
	 }

	 return res
}

func nthCatalanNumDP(n int) uint64 {
	// Create an cache
	cache := make([]uint64, n+1)

	// Init values
	cache[0] = 1
	cache[1] = 1

	// Calculate the remaining and add to cache as you go
	for i := 2; i <= n; i++ {
		//fmt.Printf("-------------i = %d--------------\n", i)
		for j := 0; j < i; j++ {
			cache[i] += cache[j]*cache[i-1-j]
			//fmt.Printf("Calculated: C[%d] * C[%d] = %d\n", j, (i-1-j), (cache[j]*cache[i-1-j]))
		}
		//fmt.Printf("--------------------------------\n")
	}

	return cache[n]
}

func main() {
	fmt.Println("find the nth catalan number")
	n := 4

	var start time.Time
	var end time.Duration

	start = time.Now()
	cNum := nthCatalanNumber(n)
	end = time.Since(start)
	fmt.Printf("Recursive: C%d = %v\n", n, cNum)
	fmt.Printf("Time to calculate recursively: %vns\n\n", end.Nanoseconds())


	start = time.Now()
	cDPNum := nthCatalanNumDP(n)
	end = time.Since(start)
	fmt.Printf("DP: C%d = %v\n", n, cDPNum)
	fmt.Printf("Time to calculate using DP: %vns\n", end.Nanoseconds())
}