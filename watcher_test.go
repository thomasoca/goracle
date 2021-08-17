package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/fsnotify/fsnotify"
)

func TestTypeConverter(t *testing.T) {
	tests := []struct {
		name   string
		event  string
		expect fsnotify.Op
	}{
		{name: "Test create event", event: "create", expect: fsnotify.Create},
		{name: "Test write event", event: "write", expect: fsnotify.Write},
		{name: "Test remove event", event: "remove", expect: fsnotify.Remove},
		{name: "Test rename event", event: "rename", expect: fsnotify.Rename},
		{name: "Test chmod event", event: "chmod", expect: fsnotify.Chmod},
		{name: "Test invalid event", event: "created", expect: fsnotify.Create},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := eventTypeConverter(tt.event)
			if result != tt.expect {
				t.Errorf("eventTypeConverter() return invalid value, expect  %v, got %v", result, tt.expect)
			}
		})
	}
}

func TestExecuteCommand(t *testing.T) {
	localPath, _ := os.Getwd()
	testPath := filepath.Join(localPath, "test", "test.txt")
	testInput := CliInput{directory: testPath, pattern: "*", eventType: "create", commandArgs: []string{"touch"}, nonBlocking: false}
	result := executeCommand(testInput, testPath)
	// Test the executeCommand function
	if result != nil {
		t.Errorf("executeCommand returned %v, want %v", result, nil)
	}
	_, err := os.Stat(testPath)
	// Test whether the command is executed
	if err != nil {
		t.Errorf("executeCommand failed to run the given command")
	}
	_ = os.Remove(testPath)
}

func TestExecuteCommandMultipleArgs(t *testing.T) {
	localPath, _ := os.Getwd()
	testPath := filepath.Join(localPath, "test", "test.txt")
	testPath2 := filepath.Join(localPath, "test", "test_2.txt")
	args := []string{"touch", testPath2}
	testInput := CliInput{directory: testPath, pattern: "*", eventType: "create", commandArgs: args, nonBlocking: false}
	result := executeCommand(testInput, testPath)

	// Test the executeCommand function
	if result != nil {
		t.Errorf("executeCommand returned %v, want %v", result, nil)
	}
	// Test whether the first file exists
	_, err := os.Stat(testPath)
	if err != nil {
		t.Errorf("executeCommand failed to create the first file")
	}
	_ = os.Remove(testPath)

	// Test whether the second file exists
	_, err = os.Stat(testPath2)
	if err != nil {
		t.Errorf("executeCommand failed to run the second file")
	}
	_ = os.Remove(testPath2)
}
