package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)


func handleArgs(args []string) (path string, recursive bool) {
	// Exists with a status code of 1 on parsing error with a message
	var err error = nil

	if len(args) != 2 {
		fmt.Println("Not enough args")
		os.Exit(1)
	}

	// Like python's os.path.exists(path)
	if _, err = os.Stat(args[0]); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s\n", args[0])
		//fmt.Println(err)
		os.Exit(1)
	} else {
		path = args[0]
	}

	if args[1] == "true" {
		recursive = true
	} else if args[1] == "false" {
		recursive = false
	} else {
		fmt.Println("Second argument must be 'true' or 'false'")
		os.Exit(1)
	}

	return path, recursive
}


type Command struct {
	// some GUI/CLI for users to add/remove paths from watch
	path string
	// should status codes be used or should the function be passed?
	exitP bool
}


func main() {
	path, recursive := handleArgs(os.Args[1:])
	// TODO: get these values from somewhere
	excludes := []string{"/home/gaige/Dropbox/school"}
	excludes = CollectExcludes(excludes)

	manager := make(chan *Command)

	watcher, watchCount, err := StartWatch(path, recursive, excludes)
	if err != nil {
		fmt.Println("Error with watcher, main.go line 59:", err)
	}

	fmt.Println("\nDirectories watched: ", watchCount, "\n")

	go EventHandler(watcher, manager)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	for {
		select {
		case sig := <-c:
			com := Command{"", true}
			manager <- &com
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("Got Signal: ", sig)
			return
		}
	}
}
