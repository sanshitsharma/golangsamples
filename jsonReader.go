package main

import (
	"fmt"
	"reflect"
)

func main() {
	//arr := []interface{}{"gig0/0", "gig0/1", "gig0/2"}
	//arr := []interface{}{1, 2, 3}
	arr := []interface{}{1.123, 2.234, 3.345}

	if reflect.TypeOf(arr).Kind() ==  reflect.Slice {
		fmt.Println("It's an array")

		switch arr[0].(type) {
		case int:
			fmt.Println("array of integers")
		case string:
			fmt.Println("array of strings")
		case float64:
			fmt.Println("array of floats")
		default:
			fmt.Println("unknown type")
		}
	}
}
