package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

type LogLevel string

const (
	LogLevelInfo    LogLevel = "INFO"
	LogLevelTrace   LogLevel = "TRACE"
	LogLevelWarning LogLevel = "WARNING"
	LogLevelError   LogLevel = "ERROR"
)

type Log struct {
	time    *time.Time
	level   LogLevel
	event   string
	message string
}

func main() {

	// Readfile loads data fully. can be used for small files
	f, err := os.ReadFile("example.log")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(f))

	lines := strings.Split(string(f), "\n")
	fmt.Println(lines[0])
	fmt.Println("Number fo lines: ", len(lines))

	fmt.Println(time.Now())

	layout := "2006-01-02 15:04:05"

	for _, line := range lines {
		// fmt.Println(i, ": ", line)

		//parts := strings.SplitN(line, " ", 5) // consecutive spaces is considered as a new one

		re := regexp.MustCompile(`\s+`)
		parts := re.Split(line, 5)
		ts, err := time.Parse(layout, "2025-"+parts[0][:2]+"-"+parts[0][3:]+" "+parts[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("time", ts)
		// fmt.Println(i, parts[0], "||", parts[1], "||", parts[2], "||", parts[3], "||", parts[4])
	}
}
