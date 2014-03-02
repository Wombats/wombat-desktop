package main

import (
	"fmt"
	"os"
	"path/filepath"
	"code.google.com/p/go.exp/fsnotify"
)


func AddWatch(path string, recursive bool, manager chan *Command) {
	//TODO: Check and handle a non-recursive watch request

	watched := 0
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
	}

	err = filepath.Walk(path, func (path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			go func (path string) (err error) {
				err = watcher.Watch(path)
				if err != nil {
					fmt.Println(err)
					return err
				}
				return nil
				// increment how many dirs are watched
				//watched++
			}(path)
		}
		return err
	})
	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case ev := <-watcher.Event:
			fmt.Println(ev)
		case err := <- watcher.Error:
			fmt.Println(err)
		}
	}
}
