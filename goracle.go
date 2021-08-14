package main

import (
	"fmt"
	"os"
)

func exitApp(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	fmt.Printf("Please use\n %s --help \nto see the CLI options and arguments\n", os.Args[0])
	os.Exit(1)
}

func main() {
	displayHelp()
	input, err := cliParser()
	if err != nil {
		exitApp(err)
	}

	goracle(input)
}
