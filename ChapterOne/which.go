package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("Please provide an arguments!\n")
		return
	}
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for i := 1; i < len(arguments); i++ {

		if i == 0 {
			continue
		}
		file := arguments[i]
		data := []string{}
		for _, directory := range pathSplit {
			fullPath := filepath.Join(directory, file)
			fileInfo, err := os.Stat(fullPath)
			if err == nil {
				mode := fileInfo.Mode()
				if mode.IsRegular() {
					if mode&0111 != 0 {
						data = append(data, fullPath)
					}
				}
			}

		}

		if len(data) == 0 {
			fmt.Printf("No search results for %s\n", file)
			continue
		}
		fmt.Printf("search results for %s:\n", file)
		for _, d := range data {
			fmt.Println(d)
		}
	}

	return
}
