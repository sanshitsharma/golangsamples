package main

import (
	"fmt"
	"io/ioutil"
)

func getDirs(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	dirs := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		} else {
			//fmt.Println(file.Name(), " is not a directory.. IGNORING")
		}
	}

	return dirs, nil
}

func findFiles(dirPath, dirId string) error {
	dirs, err := getDirs(dirPath)
	if err != nil {
		return err
	}

	if len(dirs) == 0 { // Termination condition
		// Evaluate and populate the alerts here
		fmt.Printf("DISCOVERED.. dirPath: %v, dirId: %v\n", dirPath, dirId)
		return nil
	}

	getAtId := func(dir string) string {
		if dirId == "" {
			return dir
		} else {
			return dirId + "_" + dir
		}
	}

	for _, dir := range dirs {
		findFiles(dirPath + "/" + dir, getAtId(dir))
	}

	return nil
}

func load(basePath string) error {
	return findFiles(basePath, "")
}

func main() {
	fmt.Println("finding files")
	load(".")
}
