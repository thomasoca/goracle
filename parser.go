package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func isValidevent(event string) bool {
	switch event {
	case
		"create",
		"write",
		"remove",
		"rename",
		"chmod":
		return true
	}
	return false
}

func cliParser() (CliInput, error) {
	if len(os.Args) < 2 {
		return CliInput{}, errors.New("command line argument is required to do something")
	}

	localPath, err := os.Getwd()
	if err != nil {
		return CliInput{}, err
	}
	directory := flag.String("dir", localPath, "Directory to watch, set to current working directory as default")
	pattern := flag.String("pattern", "*", "File pattern to notify")
	event := flag.String("e", "create", "Event to notify, select between create, write, remove, rename, chmod")
	nonBlocking := flag.Bool("nb", false, "Set execution mode to non-blocking")
	flag.Parse()
	cmdArguments := flag.Args()
	eventValidation := isValidevent(*event)
	if !eventValidation {
		return CliInput{}, errors.New("select only between create, write, remove, rename, chmod")
	}
	return CliInput{directory: *directory, pattern: *pattern, eventType: *event, commandArgs: cmdArguments, nonBlocking: *nonBlocking}, nil
}

func displayHelp() {
	flag.Usage = func() {
		fmt.Printf("#### goracle ####\nA CLI app to watch a directory and doing something, powered by golang and fsnotify\n")
		fmt.Printf("Usage: %s [options] COMMAND ARGS\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Printf("NOTE: Separate multiple ARGS by space\n")
		fmt.Printf("Example:\nrun a python program over a write event of .csv file\n")
		fmt.Printf("goracle -e write -dir /home/users/folder -nb -pattern *.csv python example.py input_argument\n")
		fmt.Printf("The event file name will always be passed to the last ARGS by default\n")
	}
}
