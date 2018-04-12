package main

import (
	"fmt"
)

func isParantheses(char rune) bool {
	return char == '(' || char == ')'
}

func isBalanced(str string) bool {
	counter := 0

	for _, char := range str {
		if char == '(' {
			counter++
		} else if char == ')' {
			counter--
		}

		if counter < 0 {
			return false
		}
	}

	return counter == 0
}

func main() {
	str := `()())()`
	fmt.Println(isBalanced(str))

}
