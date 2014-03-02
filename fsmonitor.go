package main

import (
	"fmt"
	"os"
	"path/filepath"
	"code.google.com/p/go.exp/fsnotify"
)


func StartWatch(path string, recursive bool) (*fsnotify.Watcher, int, error) {
	//TODO: Check and handle a non-recursive watch request

	watched := 0
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error with establishing watcher, fsmonitor.go line 17:", err)
	}

	err = filepath.Walk(path, func (path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			go func (path string) (err error) {
				err = watcher.Watch(path)
				if err != nil {
					fmt.Printf("fsmonitor.go line 25\terror: %v: %v\n", err, path)
					return err
				}
				watched++
				return nil
				// increment how many dirs are watched
			}(path)
		}
		return err
	})
	if err != nil {
		fmt.Println("Error with walking filepath, fsmonitor.go line 36:", err)
	}

	return watcher, watched, err
}


func EventHandler(watcher *fsnotify.Watcher, manager chan *Command) {
	for {
		select {
		case ev := <-watcher.Event:
			//encrypt() upload()\
			if ev.IsCreate() {
				fmt.Println("Create: ", ev)
			} else if ev.IsDelete() {
				fmt.Println("Delete: ", ev)
			} else if ev.IsModify() {
				fmt.Println("Modify: ", ev)
			} else if ev.IsRename() {
				fmt.Println("Rename: ", ev)
			} else if ev.IsAttrib() {
				fmt.Println("Attrib: ", ev)
			} else {
				fmt.Println("Event but not type?")
			}
		case err := <- watcher.Error:
			fmt.Println(err)
		case com := <-manager:
			if com.exitP {
				err := watcher.Close()
				fmt.Println("Returning EventHandler")
				if err != nil {
					fmt.Println("Error on close of watch: ", err)
				}
				return
			}
		}
	}
}
