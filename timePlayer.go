package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05.00 +0000 UTC"

	t := time.Now()
	timeStr := t.Format(layout)
	fmt.Println(timeStr)

	t1, _ := time.Parse(layout, timeStr)
	fmt.Println(t1)
}
