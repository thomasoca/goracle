package main

import (
	"fmt"
	"os"

	app "github.com/thomasoca/goracle/app"
)

func exitApp(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	fmt.Printf("Please use\n %s --help \nto see the CLI options and arguments\n", os.Args[0])
	os.Exit(1)
}

func main() {
	app.DisplayHelp()
	input, err := app.CliParser()
	if err != nil {
		exitApp(err)
	}

	app.Goracle(input)
}
