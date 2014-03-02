package main

import (
	"fmt"
	"os"

	"code.google.com/p/go.exp/fsnotify"
)


func handleArgs(args []string) (path string, recursive bool)  {
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
	path         string
	removeWatchP bool
	addWatchP    bool
	exitP        bool
}


func main() {

	path, recursive := handleArgs(os.Args[1:])

	manager := make(chan *Command)


	AddWatch(path, recursive, manager)

	//HandleEvents? or do this in the above and have main deal with SIGNALS and etc

}
