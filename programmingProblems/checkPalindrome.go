package main

import (
	"fmt"
)

func isPalindrome(str string) bool {
	if str == "" {
		return false
	}

	i := 0
	j := len(str) - 1

	for i <= j {
		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func main() {
	fmt.Println("Given a string, check if it's a palindrome")

	str := "abcdba"
	fmt.Println(`'` + str + `'`,"is Palindrome:", isPalindrome(str))
}
