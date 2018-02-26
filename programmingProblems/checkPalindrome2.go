package main

import (
	"fmt"
)

// Given a string, check if it's a palindrome. If not, check if it can be converted to palindrome
// by removing at most 1 character
func canBePalindrome(str string) bool {
	if str == "" {
		return false
	}

	i := 0
	j := len(str) - 1

	removeCount := 0
	var removedChar byte

	for i <= j {
		if str[i] != str[j] {
			if str[i+1] == str[j] {
				removeCount += 1
				removedChar = str[i]
				i += 1
				continue
			} else if str[i] == str[j-1] {
				removeCount += 1
				removedChar = str[j]
				j -= 1
				continue
			} else {
				fmt.Println(`'`+str+`'`, "cannot be converted to palindrome by removing 1 character")
				return false
			}
		}

		if removeCount > 1 {
			fmt.Println(`'`+str+`'`, "cannot be converted to palindrome by removing 1 character. Count > 1")
			return false
		}

		i++
		j--
	}

	if removeCount == 1 {
		fmt.Println(`'`+str+`'`, `can be converted to palindrome by removing`, `'`+string(removedChar)+`'`)
	} else {
		fmt.Println(`'`+str+`'`, `is already a palindrome`)
	}
	return true
}

func main() {
	str := "abdcba"
	canBePalindrome(str)
}
