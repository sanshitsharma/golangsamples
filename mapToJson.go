package main

import (
	"fmt"
	"encoding/json"
)

type DevDetails struct {
	Uuid   string	`json:"uuid"`
	Ipaddr string	`json:"ipaddr"`
}

func main() {
	device := make(map[string]DevDetails)
	device["home"] = DevDetails{
		Uuid:   "abcdef-ghijklm",
		Ipaddr: "10.68.23.11",
	}

	fmt.Println("device --> ", device)

	//Marshal to JSON
	jsonText, err := json.Marshal(device)

	fmt.Println("JSON Text: ", string(jsonText))
	fmt.Println("error: ", err)
}
