package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"strings"
)

func calcSha(dir string) {
	hasher := sha256.New()
	s, _ := ioutil.ReadDir(dir)

	for _, fileInfo := range s {
		if strings.Contains(fileInfo.Name(), ".go") {
			fileName := dir + "/" + fileInfo.Name()
			file, _ := ioutil.ReadFile(fileName)

			hasher.Write(file)
		}

		log.Println(fileInfo.Name(), ":", hex.EncodeToString(hasher.Sum(nil)))
	}

	/*
		hasher.Write(s)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(hex.EncodeToString(hasher.Sum(nil)))
	*/
}

func varArgs(names ...string) {
	if names == nil {
		log.Println("NIL args")
	} else {
		log.Println("Len args:", len(names))
		log.Println("Args:", names[0])
	}
}

func main() {
	log.Println("Welcome to checksum calculator..")

	/*
		dir := "/Users/sansshar/Desktop/test"
		calcSha(dir)
	*/

	//fpHeaders := "calcSha256.go"

	calcSha(".")
}
