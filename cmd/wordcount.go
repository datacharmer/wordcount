package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Counters struct {
	Words   bool
	Bytes   bool
	Chars   bool
	Lines   bool
	Spaces  bool
	Lower   bool
	Upper   bool
	LogFile string
}

var Version = getVersion()
var showVersion = false

func RunMain() int {
	var counters Counters

	// Same options used by `wc`
	flag.BoolVar(&counters.Lines, "l", counters.Lines, "shows number of lines")
	flag.BoolVar(&counters.Words, "w", counters.Words, "shows number of words")
	flag.BoolVar(&counters.Chars, "m", counters.Chars, "shows number of characters")
	flag.BoolVar(&counters.Bytes, "c", counters.Bytes, "shows number of bytes")

	// Extra options for this program
	flag.BoolVar(&counters.Spaces, "s", counters.Spaces, "shows number of spaces")
	flag.BoolVar(&counters.Lower, "o", counters.Lower, "shows number of lowercase characters")
	flag.BoolVar(&counters.Upper, "u", counters.Upper, "shows number of uppercase characters")
	flag.StringVar(&counters.LogFile, "log-file", counters.LogFile, "writes log file")
	flag.BoolVar(&showVersion, "version", showVersion, "shows version")
	flag.Parse()

	if showVersion {
		fmt.Println(Version)
		return 0
	}
	err := runWordCount(counters)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return 1
	}
	return 0
}

func runWordCount(counters Counters) error {
	scanner := bufio.NewScanner(os.Stdin)
	var logFile *os.File
	var err error
	if counters.LogFile != "" {
		logFile, err = os.Create(counters.LogFile)
		if err != nil {
			return fmt.Errorf("error creating file '%s': %s", counters.LogFile, err)
		}
		defer logFile.Close()
	}
	var lines []string
	var totalBytes int
	var totalChars int
	var totalWords int
	var totalSpaces int
	var totalLower int
	var totalUpper int
	if err = writeLog(logFile, fmt.Sprintf("wanted lines:     %v\n", counters.Lines)); err != nil {
		return err
	}
	if err = writeLog(logFile, fmt.Sprintf("wanted bytes:     %v\n", counters.Bytes)); err != nil {
		return err
	}
	if err = writeLog(logFile, fmt.Sprintf("wanted chars:     %v\n", counters.Chars)); err != nil {
		return err
	}
	if err = writeLog(logFile, fmt.Sprintf("wanted words :    %v\n", counters.Words)); err != nil {
		return err
	}
	if err = writeLog(logFile, fmt.Sprintf("wanted spaces:    %v\n", counters.Spaces)); err != nil {
		return err
	}
	if err = writeLog(logFile, fmt.Sprintf("wanted lowercase: %v\n", counters.Lower)); err != nil {
		return err
	}
	if err = writeLog(logFile, fmt.Sprintf("wanted uppercase: %v\n", counters.Upper)); err != nil {
		return err
	}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		if err = writeLog(logFile, fmt.Sprintf("processing line %3d: %s\n", len(lines), line)); err != nil {
			return err
		}

		totalBytes += len(line)
		totalChars += len([]rune(line))
		totalWords += len(strings.Fields(line))
		for _, r := range line {
			if unicode.IsLower(r) {
				totalLower++
			}
			if unicode.IsUpper(r) {
				totalUpper++
			}
			if unicode.IsSpace(r) {
				totalSpaces++
			}
		}
	}

	numLines := len(lines)

	anyOption := counters.Words || counters.Chars || counters.Bytes || counters.Lines ||
		counters.Spaces || counters.Lower || counters.Upper
	if !anyOption {
		counters.Words = true
		counters.Bytes = true
		counters.Lines = true
	}
	if counters.Spaces {
		fmt.Println(totalSpaces)
		return nil
	}
	if counters.Lower {
		fmt.Println(totalLower)
		return nil
	}
	if counters.Upper {
		fmt.Println(totalUpper)
		return nil
	}
	if err = writeLog(logFile, fmt.Sprintf("lines found: %d\n", numLines)); err != nil {
		return err
	}
	if err = writeLog(logFile, fmt.Sprintf("words found: %d\n", totalWords)); err != nil {
		return err
	}
	if err = writeLog(logFile, fmt.Sprintf("bytes found: %d\n", totalBytes)); err != nil {
		return err
	}
	if err = writeLog(logFile, fmt.Sprintf("chars found: %d\n", totalChars)); err != nil {
		return err
	}
	if err = writeLog(logFile, fmt.Sprintf("spaces found: %d\n", totalSpaces)); err != nil {
		return err
	}
	if err = writeLog(logFile, fmt.Sprintf("lowercase found: %d\n", totalLower)); err != nil {
		return err
	}
	if err = writeLog(logFile, fmt.Sprintf("uppercase found: %d\n", totalUpper)); err != nil {
		return err
	}
	previous := ""
	if counters.Lines {
		fmt.Printf("%d", numLines)
		previous = " "
	}
	if counters.Words {
		fmt.Printf("%s%d", previous, totalWords)
		previous = " "
	}
	if counters.Chars {
		fmt.Printf("%s%d", previous, totalChars)
		previous = " "
		counters.Bytes = false
	}
	if counters.Bytes {
		fmt.Printf("%s%d", previous, totalBytes)
	}
	fmt.Println()
	return nil
}

func writeLog(logFile *os.File, item string) error {
	if logFile == nil {
		return nil
	}
	_, err := logFile.WriteString(item)
	if err != nil {
		return err
	}
	return nil
}

func getVersion() string {
	version := os.Getenv("WCVERSION")
	if version == "" {
		version = "0.2"
	}
	return version

}
