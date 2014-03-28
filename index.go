package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)


type Path struct {
	Path string
	IsDir bool
	ModTime time.Time
}


func IndexAndCompare(paths []string, excludes []string) {
	generateLocalIndex(paths)
}


func generateLocalIndex(paths []string) {
	index := make([]Path, 1, 2)
	for _, path := range paths {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			index = append(index, info)
			fmt.Println(info.ModTime())
			return err
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}
