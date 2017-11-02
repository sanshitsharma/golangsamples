package main

import (
	"encoding/json"
	"fmt"
)

type myStruct struct {
	Name 	     string `json:"name"`
	Id 	     uint32 `json:"id"`
	IsRegistered bool `json:"is_registered"`
	Phones       []string
}

func main() {
	myInst := &myStruct{
		Name: "Sanshit Sharma",
		Id: 1806,
		IsRegistered: true,
		Phones: []string{"513-549-1487", "408-250-5749"},
	}

	myInstBytes, err := json.Marshal(myInst)
	if err != nil {
		fmt.Printf("error marshalling data %v\n", err)
	} else {
		fmt.Printf("myInstString: '%v'\n", string(myInstBytes))
	}
}

