package main

import "fmt"

func isSafe(maze [][]int, x, y int) bool {
	if x < len(maze) && y < len(maze[0]) && maze[x][y] == 1 {
		return true
	}

	return false
}

func traverseMazeUtil(maze [][]int, x, y int, sol [][]int) bool {
	if x == len(maze)-1 && y == len(maze[0])-1 {
		sol[x][y] = 1
		return true
	}

	if isSafe(maze, x, y) {

		sol[x][y] = 1

		if traverseMazeUtil(maze, x+1, y, sol) {
			return true
		}

		if traverseMazeUtil(maze, x, y+1, sol) {
			return true
		}

		sol[x][y] = 0
		return false
	}

	return false
}

func traverseMaze(maze [][]int) bool {
	rows := len(maze)
	cols := len(maze[0])

	sol := make([][]int, rows)
	for i := range sol {
		sol[i] = make([]int, cols)
	}

	// Init conditions
	x := 0
	y := 0
	sol[x][y] = 1

	if traverseMazeUtil(maze, x, y, sol) {
		fmt.Println(sol)
		return true
	}

	fmt.Println("No solution found!")
	return false
}

func main() {
	probStatement := `Given a 2D matrix marked with 0s & 1s representing a maze. Check if there exists a way for a rat to reach from top left to bottom right corener`
	fmt.Println(probStatement)

	maze := [][]int{{1, 0, 0, 0}, {1, 1, 0, 1}, {0, 1, 0, 0}, {1, 1, 1, 1}}
	traverseMaze(maze)
}
