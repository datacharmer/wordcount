package main

import (
	"path"
	"strconv"
	"time"

	"github.com/rogpeppe/go-internal/testscript"
)

// sleepFor is a testscript command that pauses the execution for the required number of seconds
func sleepFor(ts *testscript.TestScript, neg bool, args []string) {
	duration := 0
	var err error
	if len(args) == 0 {
		duration = 1
	} else {
		duration, err = strconv.Atoi(args[0])
		ts.Check(err)
	}
	time.Sleep(time.Duration(duration) * time.Second)
}

// checkFile is a testscript command that checks the existence of a list of files
// inside a directory
func checkFiles(ts *testscript.TestScript, neg bool, args []string) {
	if len(args) < 1 {
		ts.Fatalf("syntax: check_file directory_name file_name [file_name ...]")
	}
	dir := args[0]

	for i := 1; i < len(args); i++ {
		f := path.Join(dir, args[i])
		if neg {
			if fileExists(f) {
				ts.Fatalf("file %s found", f)
			}
		}
		if !fileExists(f) {
			ts.Fatalf("file not found %s", f)
		}
	}
}

func customCommands() map[string]func(ts *testscript.TestScript, neg bool, args []string) {
	return map[string]func(ts *testscript.TestScript, neg bool, args []string){

		// check_files will check that a given list of files exists
		// invoke as "check_files workdir file1 [file2 [file3 [file4]]]"
		// The command can be negated, i.e. it will succeed if the given files do not exist
		// "! check_files workdir file1 [file2 [file3 [file4]]]"
		"check_files": checkFiles,

		// sleep_for will pause execution for the required number of seconds
		// Invoke as "sleep_for 3"
		// If no number is passed, it pauses for 1 second
		"sleep_for": sleepFor,
	}
}
