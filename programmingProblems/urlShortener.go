package main

import (
	"fmt"
	"math"
)

func decimalToBase62(num uint64) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	charMap := []rune(str)

	res := make([]rune, 0)

	for num != 0 {
		res = append([]rune{charMap[num%62]}, res...)
		num = num / 62
	}

	return string(res)
}

func base62ToDecimal(str string) uint64 {
	runes := []rune(str)
	var num uint64 = 0

	runes = reverse(runes)

	for indx, char := range runes {
		baseFactor := int(math.Pow(float64(62), float64(indx)))
		if char >= 'A' && char <= 'Z' {
			num += uint64(int(char-'A') * baseFactor)
		} else if char >= 'a' && char <= 'z' {
			num += uint64((int(char-'a') + 26) * baseFactor)

		} else if char >= '0' && char <= '9' {
			num += uint64((int(char-'0') + 52) * baseFactor)
		} else {
			fmt.Println("invalid rune")
		}
	}

	return num
}

func main() {
	fmt.Println("Write a URL shortner")

	var num uint64 = 18446744073709551615
	base62Str := decimalToBase62(num)
	fmt.Printf("'%d' converted to Base62 = '%s'\n", num, base62Str)

	fmt.Printf("'%s' converted to Base10 = %d\n", base62Str, base62ToDecimal(base62Str))
}

func reverse(runes []rune) []rune {
	l := 0
	h := len(runes) - 1

	for l < h {
		temp := runes[l]
		runes[l] = runes[h]
		runes[h] = temp

		l++
		h--
	}

	return runes
}
