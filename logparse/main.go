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

func logLevelParser(level string) (LogLevel, error) {
	switch strings.ToUpper(level) {
	case "INFO":
		return LogLevelInfo, nil
	case "ERROR":
		return LogLevelError, nil
	case "WARNING":
		return LogLevelWarning, nil
	case "TRACE":
		return LogLevelTrace, nil
	default:
		return "", fmt.Errorf("Unknown log level: %s", level)
	}

}

type Log struct {
	time    time.Time
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

	// this date is fixed one for layout. we also have big one as well but always use below exadct date for layout (but choose ur layout like date or time comes front etc)
	layout := "2006-01-02 15:04:05"

	logs := make([]Log, 0)
	for _, line := range lines {
		// fmt.Println(i, ": ", line)

		//parts := strings.SplitN(line, " ", 5) // consecutive spaces is considered as a new one

		// remove ending empty line case
		if strings.TrimSpace(line) == "" {
			continue
		}

		re := regexp.MustCompile(`\s+`)
		parts := re.Split(line, 5)
		ts, err := time.Parse(layout, "2025-"+parts[0][:2]+"-"+parts[0][3:]+" "+parts[1])
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println("my parts", parts)

		loglevel, err := logLevelParser(parts[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ts)
		logs = append(logs, Log{ts, loglevel, parts[3], parts[4]})

		// fmt.Println("time", ts)
		// fmt.Println(i, parts[0], "||", parts[1], "||", parts[2], "||", parts[3], "||", parts[4])
	}
	log.Println(logs)
}
