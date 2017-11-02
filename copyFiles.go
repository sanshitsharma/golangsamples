package main

import (
	"log"
	"os/exec"
)

func passByRef(a []string) {
	a = append(a, "Sanshit")
	a = append(a, "Sharma")

	log.Println(a)
}

func main() {
	log.Println("Welcome to copier")
	//createMopsDir()

	arr := make([]string, 0)
	arr = append(arr, "Vrishti")

	a := make([]string, 0)
	a = append(a, "Sanshit")
	a = append(a, "Sharma")

	arr = append(arr, a...)
	log.Println("Array: ", arr)
}