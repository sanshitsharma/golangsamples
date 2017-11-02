package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
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
				 Name    string            `json:"name"`
				 Columns []string          `json:"columns"`
				 Values  [][]interface{}   `json:"values"`
			 } `json:"series"`
		 } `json:"data"`
}

func main() {
	fmt.Println("Welcome to Json decoder")

	// read a json string
	dat, err := ioutil.ReadFile("/Users/sansshar/go/src/sansshar_progs/JsonDecoder/sample.json")
	if err != nil {
		fmt.Println("Could not read file")
	}

	fmt.Printf("Alert\n%s", dat)

	payload := tickAlert{}
	err = json.Unmarshal([]byte(dat), &payload)

	fmt.Printf("Parsed data\n%v", payload.ID)
}