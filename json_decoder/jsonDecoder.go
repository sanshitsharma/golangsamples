package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type tickAlert struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	Details  string `json:"details"`
	Time     string `json:"time"`
	Duration uint64 `json:"duration"`
	Level    string `json:"level"`
	Data     struct {
		Series []struct {
			Name    string          `json:"name"`
			Columns []string        `json:"columns"`
			Values  [][]interface{} `json:"values"`
		} `json:"series"`
	} `json:"data"`
}

func main() {
	fmt.Println("Welcome to Json decoder")

	// read a json string
	dat, err := ioutil.ReadFile("sample.json")
	if err != nil {
		fmt.Println("Could not read file")
	}

	fmt.Println("Alert:", string(dat))

	payload := tickAlert{}
	err = json.Unmarshal([]byte(dat), &payload)

	fmt.Println("Parsed data:", payload.ID)
}
