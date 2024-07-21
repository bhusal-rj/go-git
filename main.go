package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: git-go <commands> [<args>]")
		os.Exit(1)
	}

	command := os.Args[0]
	args := os.Args[2:]

	currentWorkDir, err := os.Getwd()
	if err != nil {
		fmt.Println("There is an error geting current dir", err)
		os.Exit(1)
	}
	fmt.Sprintf("%v %v %v", command, args, currentWorkDir)
}
