package main

import (

	"fmt"
	"os"
	"path/filepath"
	"strings"
)


func CollectExcludes(excludePaths []string) ([]string){

	excludes := make([]string, 1, 1)

	for i := 0; i < len(excludePaths); i++ {
		exclude := excludePaths[i]
		err := filepath.Walk(exclude, func(path string, info os.FileInfo, err error) (error) {
			if (strings.HasPrefix(path, exclude)) {
				excludes = append(excludes, path)
			}
			return err
		})

		if err != nil {
			fmt.Println("Error collecting excludes: ", err)
		}
	}
	return excludes
}


func HasMember(arr []string, item string) (bool) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			return true
		}
	}
	return false
}
