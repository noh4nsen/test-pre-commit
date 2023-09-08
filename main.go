package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			if err := checkFile(path); err != nil {
				fmt.Printf("Error in %s: %s\n", path, err)
				os.Exit(1)
			}
		}
		return nil
	})
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
