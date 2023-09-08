package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: test-pre-commit <file_or_directory>")
		os.Exit(1)
	}

	target := os.Args[1]
	fileInfo, err := os.Stat(target)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if fileInfo.IsDir() {
		filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
				if err := checkFile(path); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
			return nil
		})
	} else {
		if err := checkFile(target); err != nil {
			fmt.Printf("Error in %s: %s\n", target, err)
			os.Exit(1)
		}
	}
}

func checkFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if strings.Contains(string(content), "TODO") {
		return fmt.Errorf("found 'TODO' in %s", filename)
	}

	return nil
}
