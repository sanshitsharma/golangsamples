package main

import (
	"fmt"
	"reflect"
)

func reflector(payload string) {
	fmt.Printf("From Method\n%v", reflect.TypeOf(payload))
}

func main() {
	// Read in the file

	payload := `{ "id": "` + "1499301099498-0ddfd012-4fd1-48f1-b465-4303e0198bf6" +
		`", "task": "` + "xyz" +
		`", "result": "` + "SUCCESS" +
		`", "task_id": "` + "1" +
		`", "description": "` + "success/failure criteria met" + `" }`

	fmt.Println(reflect.TypeOf(payload))

	reflector(payload)
}
