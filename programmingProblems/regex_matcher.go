package main

import (
	"fmt"
)

func isMatch(s, r string) bool {
	i := 0 // s index
	j := 0 // r index

	for j < len(r) && i < len(s) {
		//fmt.Println("s[i] =", string(s[i]), " r[j] =", string(r[j]) )

		// Read token from r
		token := r[j]
		var nextChar byte
		isStar := false

		if j < len(r) - 1 && r[j+1] == '*' {
			isStar = true
		}

		if token == '.' && isStar && j < len(r) - 2 {
			nextChar = r[j+2]
		}

		if isStar {
			j += 2
		} else {
			j += 1
		}

		//fmt.Println("token:", string(token), " isStar:", isStar, " nextChar:", nextChar)

		// Now see if we can find a match for the token in s
		if token == '.' {
			if isStar {
				if nextChar == 0 {
					return true
				}

				for i < len(s) && s[i] != nextChar {
					i += 1
				}
				if i == len(s) {
					return true
				}
			} else {
				i += 1
			}
		} else { // token != '.'
			if isStar && s[i] != token {
				//fmt.Println("CONTINUE.. token:", string(token), " isStar:", isStar)
				continue
			}

			if isStar {
				//fmt.Println("finding multiple", string(token))
				for i < len(s) && s[i] == token {
					i += 1
				}

				if i == len(s) {
					return true
				}
			} else {
				i += 1
			}
		}
	}


	if i == len(s) && j == len(r) {
		return true
	}

	return false

}

func main() {
	fmt.Println(`Implement regular expression matching with support for '.' and '*'`)
	res := isMatch(`aab`, "...")

	fmt.Println(res)
}
