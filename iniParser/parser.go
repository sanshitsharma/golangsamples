package main

import (
	"fmt"
	"os"

	"github.com/dlintw/goconf"
)

func main() {
	fmt.Println("parsing ini file")

	c, err := goconf.ReadConfigFile("iniParser/sample.ini")
	if err != nil {
		fmt.Println("failed to read conf file:", err)
		os.Exit(1)

	}

	// Fetch the log file for proc2
	proc2OutLog, _ := c.GetString("program:proc2", "stdout_logfile")
	fmt.Println("Process 2 out log:", proc2OutLog)
}
