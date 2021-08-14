package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

type CliInput struct {
	directory   string
	pattern     string
	eventType   string
	commandArgs []string
	nonBlocking bool
}

func eventTypeConverter(event string) (fsnotify.Op, error) {
	var eventType fsnotify.Op
	switch event {
	case "create":
		eventType = fsnotify.Create
	case "write":
		eventType = fsnotify.Write
	case "remove":
		eventType = fsnotify.Remove
	case "rename":
		eventType = fsnotify.Rename
	case "chmod":
		eventType = fsnotify.Chmod
	default:
		return eventType, errors.New("failed to convert event string")
	}
	return eventType, nil
}

func filterDirsGlob(path, pattern string) (bool, error) {
	file := filepath.Base(path)
	return filepath.Match(pattern, file)
}

func executeCommand(cliInput CliInput, fileName string) error {
	// concat the file name to the end of the commandArgs
	commandArgs := cliInput.commandArgs
	commandArgs = append(commandArgs, fileName)

	// parse the commandArgs into golang Command function
	executable := commandArgs[0]
	commandArgs = commandArgs[1:]
	commandArgsString := strings.Join(commandArgs, " ")
	log.Println(executable, " ", commandArgsString)
	cmd := exec.Command(executable, commandArgsString)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if cliInput.nonBlocking {
		if err := cmd.Start(); err != nil {
			return err
		}
		return nil
	}
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil

}

// blatantly copied from fsnotify example
func Goracle(cliInput CliInput) {
	eventType, err := eventTypeConverter(cliInput.eventType)
	if err != nil {
		log.Fatal(err)
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fileName := event.Name
				checkFile, err := filterDirsGlob(fileName, cliInput.pattern)
				if err != nil {
					log.Fatal("Error occured: ", err)
				}

				if event.Op&eventType == eventType && checkFile {
					// do something here
					err = executeCommand(cliInput, fileName)
					if err != nil {
						log.Printf("Failed to start cmd: %v", err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(cliInput.directory)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
