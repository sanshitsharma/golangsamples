package main

import (
	"fmt"
	"runtime"
)

func MyCaller() (string, error) {
	pc := make([]uintptr, 5)
	skip := 1

	runtime.Callers(skip, pc)
	fmt.Println(pc)

	for i := 0; i < len(pc); i++ {
		fun := runtime.FuncForPC(pc[0] - 1)
		fmt.Println(fun.Name())
	}
	return "test", nil
}

func foo() {
	fmt.Println(MyCaller())
}

func bar() {
	foo()
}

func main() {
	bar()
}
