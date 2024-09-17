package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func findFiles(path string, filenameMask string) ([]string, error) {
	var matchedFiles []string

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			match, err := filepath.Match(filenameMask, info.Name())
			if err != nil {
				return err
			}
			if match {
				matchedFiles = append(matchedFiles, filePath)
			}
		} else if info.Name() == "v1" {
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return matchedFiles, nil
}

func main() {
	files, err := findFiles("D:/Projects/learning-go", "*.go")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
