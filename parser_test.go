package main

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func TestCliParser(t *testing.T) {
	localPath, _ := os.Getwd()
	// Defining our test slice. Each unit test should have the following properties:
	tests := []struct {
		name    string
		want    CliInput
		wantErr bool
		osArgs  []string
	}{
		// Test with default parameter
		{"Default parameters", CliInput{directory: localPath, pattern: "*", eventType: "create", commandArgs: []string{"echo", "hello world"}, nonBlocking: false}, false, []string{"cmd", "echo", "hello world"}},
		// Test custom directory
		{"Custom directory", CliInput{directory: "/home/projects", pattern: "*", eventType: "create", commandArgs: []string{"echo", "hello world"}, nonBlocking: false}, false, []string{"cmd", "--dir=/home/projects", "echo", "hello world"}},
		// Test custom pattern
		{"Custom pattern", CliInput{directory: localPath, pattern: "*.txt", eventType: "create", commandArgs: []string{"echo", "hello world"}, nonBlocking: false}, false, []string{"cmd", "--pattern=*.txt", "echo", "hello world"}},
		// Test custom event type
		{"Custom event", CliInput{directory: localPath, pattern: "*", eventType: "write", commandArgs: []string{"echo", "hello world"}, nonBlocking: false}, false, []string{"cmd", "--e=write", "echo", "hello world"}},
		// Test non blocking mode
		{"Non blocking", CliInput{directory: localPath, pattern: "*", eventType: "create", commandArgs: []string{"echo", "hello world"}, nonBlocking: true}, false, []string{"cmd", "-nb", "echo", "hello world"}},
		// Test missing parameter
		{"No parameters", CliInput{}, true, []string{"cmd"}},
		// Test invalid event
		{"Invalid event", CliInput{}, true, []string{"cmd", "--e=overwrite", "echo", "hello world"}},
	}
	// Iterating over the previous test slice
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualOsArgs := os.Args
			defer func() {
				os.Args = actualOsArgs
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			}()

			os.Args = tt.osArgs
			got, err := cliParser()
			if (err != nil) != tt.wantErr {
				t.Errorf("cliParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cliParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidEvent(t *testing.T) {
	tests := []struct {
		name   string
		event  string
		expect bool
	}{
		{name: "Test create event", event: "create", expect: true},
		{name: "Test write event", event: "write", expect: true},
		{name: "Test remove event", event: "remove", expect: true},
		{name: "Test rename event", event: "rename", expect: true},
		{name: "Test chmod event", event: "chmod", expect: true},
		{name: "Test invalid event", event: "created", expect: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidevent(tt.event)
			if result != tt.expect {
				t.Errorf("isValidEvent() return invalid value, expect  %v, got %v", result, tt.expect)
			}
		})
	}
}
