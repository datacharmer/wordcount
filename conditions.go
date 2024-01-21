package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/datacharmer/wordcount/cmd"
	versions "github.com/hashicorp/go-version"
)

// customConditions is a testscript function that handles all the conditions defined for this test
func customConditions(condition string) (bool, error) {
	// assumes arguments are separated by a colon (":")
	elements := strings.Split(condition, ":")
	if len(elements) == 0 {
		return false, fmt.Errorf("no condition found")
	}
	name := elements[0]
	switch name {
	case "version_is_at_least":
		return versionIsAtLeast(elements)
	case "exists_within_seconds":
		return existsWithinSeconds(elements)
	default:
		return false, fmt.Errorf("unrecognized condition '%s'", name)
	}
}

func versionGreaterOrEqual(v1, v2 string) (bool, error) {
	ver1, err := versions.NewVersion(v1)
	if err != nil {
		return false, err
	}
	ver2, err := versions.NewVersion(v2)
	if err != nil {
		return false, err
	}
	return ver1.GreaterThanOrEqual(ver2), nil
}

func versionIsAtLeast(elements []string) (bool, error) {
	if len(elements) < 2 {
		return false, fmt.Errorf("condition '%s' requires version", elements[0])
	}
	version := elements[1]
	result, err := versionGreaterOrEqual(cmd.Version, version)
	if result && os.Getenv("WCDEBUG") != "" {
		fmt.Printf("Current version '%s' is at least as requested '%s'\n", cmd.Version, version)
	}
	return result, err
}

func existsWithinSeconds(elements []string) (bool, error) {
	if len(elements) < 3 {
		return false, fmt.Errorf("condition 'exists_within_seconds' requires a file name and the number of seconds")
	}
	fileName := elements[1]
	delay, err := strconv.Atoi(elements[2])
	if err != nil {
		return false, err
	}
	if delay == 0 {
		return fileExists(fileName), nil
	}
	elapsed := 0
	for elapsed < delay {
		time.Sleep(time.Second)
		if fileExists(fileName) {
			return true, nil
		}
		elapsed++
	}
	fmt.Printf("file %s not found within %d seconds\n", fileName, delay)
	return false, nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
