package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func LoopFile(root string, ext [] string) chan string {
	files := make(chan string)

	go func() {
		err := filepath.Walk(root, visit(files, ext))
		if err != nil {
			panic(err)
		}

		defer close(files)
	}()

	return files
}

func visit(files chan string, ext []string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if info.IsDir() {
			return nil
		}

		for _, e := range ext {
			if strings.EqualFold(filepath.Ext(path), "."+e) {
				files <- path
				break
			}
		}

		return nil
	}
}
