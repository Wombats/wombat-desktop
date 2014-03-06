package main

import (
	"code.google.com/p/go.exp/fsnotify"
	"fmt"
	"os"
	"path/filepath"
)

func StartWatch(path string, recursive bool, excludes []string) (*fsnotify.Watcher, int, error) {
	// TODO: Check and handle a non-recursive watch request
	// TODO: Handle directory excludes on startup
	watched := 0
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error with establishing watcher, fsmonitor.go line 17:", err)
	}

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && !(HasMember(excludes, path)) {
			go func(path string) (err error) {
				fmt.Println(path)
				err = watcher.Watch(path)
				if err != nil {
					fmt.Printf("fsmonitor.go line 25\terror: %v: %v\n", err, path)
					return err
				}
				watched++
				return nil
				// increment how many dirs are watched
				// TODO: try to find out why the number of directories
				//       watched seems to be different between executions
			}(path)
		}
		return err
	})
	if err != nil {
		fmt.Println("Error with walking filepath, fsmonitor.go line 36:", err)
	}

	return watcher, watched, err
}

//Perhaps this later changes to logEvent or something
func logEvent(name string, eventType string) {
	// after deletion (and potentially rename) we cannot ascertain
	// if the thing renamed or deleted was a file or directory. This
	// more or may not be a problem.
	info, err := os.Lstat(name)
	if err != nil {
		fmt.Printf("File or directory %s: %v\n", eventType, name)
		return
	}
	if info.IsDir() {
		fmt.Printf("Directory %s: %v\n", eventType, name)
	} else if !(info.IsDir()) {
		fmt.Printf("File %s: %v\n", eventType, name)
	}
	return
}

func EventHandler(watcher *fsnotify.Watcher, manager chan *Command) {
	for {
		select {
		case ev := <-watcher.Event:
			//encrypt() upload()
			switch {
			case ev.IsCreate():
				watcher.Watch(ev.Name)
				logEvent(ev.Name, "create")

			case ev.IsDelete():
				watcher.RemoveWatch(ev.Name)
				logEvent(ev.Name, "delete")

			case ev.IsModify():
				logEvent(ev.Name, "modify")

			case ev.IsAttrib():
				logEvent(ev.Name, "modify attrib")

			case ev.IsRename():
				watcher.RemoveWatch(ev.Name)
				logEvent(ev.Name, "rename")

			default:
				fmt.Println("Something is weird. Event but not type?")
			}
		case err := <-watcher.Error:
			// TODO: handle errors and see why reading from this can cause a block.
			fmt.Println(err)
		case com := <-manager:
			// TODO: Add in ability to add/remove watches from a recieved command
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
