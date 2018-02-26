package main

import "fmt"

func newPosition(n, diagIndx, x, y int, isTopRight bool) (int, int) {
	numMoves := n - 1
	moved := 0

	if numMoves == diagIndx {
		return x, y
	}

	if isTopRight {
		fmt.Println("numMoves =", numMoves)
		for y < numMoves && numMoves > moved {
			y++
			moved++
		}
		for x < numMoves && numMoves > moved {
			x++
			moved++
		}
		for y >= diagIndx && numMoves > moved {
			y--
			moved++
		}
		for x >= diagIndx && numMoves > moved {
			x--
			moved++
		}
	} else { // isBottomLeft
		for y > diagIndx && numMoves > moved {
			y--
			moved++
		}
		for x > diagIndx && numMoves > moved {
			x--
			moved++
		}
		for y < n && numMoves > moved {
			y++
			moved++
		}
		for x < n && numMoves > moved {
			x++
			moved++
		}
	}

	return x, y
}

func rotate(mat [][]int, n, diagIndx int) {
	i := diagIndx
	j := diagIndx

	// Read Top
	for ; j < n-1; j++ {
		newX, newY := newPosition(n, diagIndx, i, j, true)
		fmt.Printf("Top mat[%d][%d] --> mat[%d][%d]\n", i, j, newX, newY)
	}

	// Read Right
	for ; i < n; i++ {
		newX, newY := newPosition(n, diagIndx, i, j, true)
		fmt.Printf("Right mat[%d][%d] --> mat[%d][%d]\n", i, j, newX, newY)
	}

	// Read Bottom
	i--
	j--
	for ; j > diagIndx; j-- {
		newX, newY := newPosition(n, diagIndx, i, j, false)
		fmt.Printf("Bottom mat[%d][%d] --> mat[%d][%d]\n", i, j, newX, newY)
	}

	// Read Left
	for ; i > diagIndx; i-- {
		newX, newY := newPosition(n, diagIndx, i, j, false)
		fmt.Printf("Left mat[%d][%d] --> mat[%d][%d]\n", i, j, newX, newY)
	}
}

func rotate_matrix(mat [][]int) {
	n := len(mat)
	diagIndx := 0

	for diagIndx < n {
		rotate(mat, n, diagIndx)
		n--
		diagIndx++
	}
}

func main() {
	fmt.Println("You are given an n x n 2D matrix representing an image. Rotate the image by 90 degrees (clockwise).")
	mat := [][]int{{1, 2, 3, 10}, {4, 5, 6, 20}, {7, 8, 9, 30}, {10, 11, 12, 40}}
	//mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	/*
		oldX := 1
		oldY := 0
		newX, newY := newPosition(len(mat), 0, oldX, oldY, false)
		fmt.Printf("Old Position: (%d, %d). New Position: (%d, %d)\n", oldX, oldY, newX, newY)
	*/
	rotate_matrix(mat)
}
