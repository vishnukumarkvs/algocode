package main

import (
	"bufio"
	cryptoRand "crypto/rand"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/term"
)

// Simple color/style helpers (ANSI)
const (
	clrReset  = "\x1b[0m"
	clrDim    = "\x1b[2m"
	clrGreen  = "\x1b[32m"
	clrRed    = "\x1b[31m"
	clrYellow = "\x1b[33m"

	clrClear   = "\x1b[2J"
	clrHome    = "\x1b[H"
	hideCursor = "\x1b[?25l"
	showCursor = "\x1b[?25h"
)

// State per character: 0 = untyped, 1 = correct, -1 = incorrect
type mark int8

const (
	untyped mark = 0
	ok      mark = 1
	bad     mark = -1
)

func main() {
	paragraphs := []string{
		"Typing is a skill you build with practice. Focus on accuracy first, then let speed follow naturally.",
		"Gophers love simple tools that do one thing well. Keep your code small, readable, and tested.",
		"Concurrency is not parallelism, but in Go you can use goroutines to model both cleanly and safely.",
		"Measure twice, cut once. Benchmarks and profiling often reveal surprising bottlenecks in production systems.",
		"Great software is a conversation between humans. Write code that explains itself to the next reader.",
	}

	target := pickRandom(paragraphs)
	targetRunes := []rune(target)
	status := make([]mark, len(targetRunes))

	// Put terminal into raw mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Failed to enter raw mode:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	// Ensure cursor restored even on panic/exit
	defer func() {
		fmt.Print(showCursor, clrReset)
	}()

	// Handle Ctrl+C cleanly
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		<-sigCh
		term.Restore(int(os.Stdin.Fd()), oldState)
		fmt.Print(showCursor, clrReset, "\n")
		os.Exit(0)
	}()

	reader := bufio.NewReader(os.Stdin)
	started := false
	startTime := time.Time{}
	idx := 0
	var typed, correct, errors int

	// Render initial screen
	render(targetRunes, status, idx, started, startTime, typed, correct, errors, false)

	for {
		// Quit on ESC (27)
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		if b == 27 { // ESC
			break
		}

		// Handle Ctrl+C by ourselves too (in addition to signal handler)
		if b == 3 { // ^C
			break
		}

		// Backspace: 127 (DEL) or 8
		if b == 127 || b == 8 {
			if idx > 0 {
				idx--
				if status[idx] == ok {
					correct--
					typed--
				} else if status[idx] == bad {
					errors--
					typed--
				}
				status[idx] = untyped
			}
			render(targetRunes, status, idx, started, startTime, typed, correct, errors, false)
			continue
		}

		// Ignore non-printable except newline (we’ll convert to space)
		r := rune(b)
		if b < 32 || b == 127 {
			// Convert newline/tab to space to keep flow
			if b == '\n' || b == '\r' || b == '\t' {
				r = ' '
			} else {
				// other control chars ignored
				continue
			}
		}

		// Start timer on first key
		if !started {
			started = true
			startTime = time.Now()
		}

		// If finished, any key shows summary & exits
		if idx >= len(targetRunes) {
			render(targetRunes, status, idx, started, startTime, typed, correct, errors, true)
			_, _ = reader.ReadByte() // wait one key then exit
			break
		}

		// Compare & update
		expected := targetRunes[idx]
		if r == expected {
			status[idx] = ok
			correct++
		} else {
			status[idx] = bad
			errors++
		}
		typed++
		idx++

		// If we just finished
		if idx >= len(targetRunes) {
			render(targetRunes, status, idx, started, startTime, typed, correct, errors, true)
			_, _ = reader.ReadByte() // wait one key then exit
			break
		}

		render(targetRunes, status, idx, started, startTime, typed, correct, errors, false)
	}
}

func render(target []rune, status []mark, idx int, started bool, start time.Time, typed, correct, errors int, finished bool) {
	// Clear & place cursor home
	fmt.Print(hideCursor, clrClear, clrHome)

	// Title / help
	fmt.Println("🐒  monkeytype (terminal)")
	fmt.Println(clrDim + "Type the paragraph below. Backspace to correct. Esc/Ctrl+C to quit." + clrReset)
	fmt.Println()

	// Render the paragraph with color coding
	var sb strings.Builder
	for i, r := range target {
		switch status[i] {
		case ok:
			sb.WriteString(clrGreen)
			sb.WriteRune(r)
			sb.WriteString(clrReset)
		case bad:
			sb.WriteString(clrRed)
			sb.WriteRune(r)
			sb.WriteString(clrReset)
		default:
			// caret marker at current index
			if i == idx {
				sb.WriteString(clrYellow)
				sb.WriteRune(r)
				sb.WriteString(clrReset)
			} else {
				sb.WriteRune(r)
			}
		}
	}
	fmt.Println(sb.String())
	fmt.Println()

	// Stats
	elapsed := time.Duration(0)
	if started {
		elapsed = time.Since(start)
	}
	minutes := elapsed.Minutes()
	grossWPM := 0.0
	netWPM := 0.0
	if minutes > 0 {
		grossWPM = float64(typed) / 5.0 / minutes
		netWPM = float64(correct-errors) / 5.0 / minutes
		if netWPM < 0 {
			netWPM = 0
		}
	}
	acc := 100.0
	if typed > 0 {
		acc = (float64(correct) / float64(typed)) * 100.0
	}

	fmt.Printf("%sTime:%s %s  %sGross WPM:%s %.1f  %sNet WPM:%s %.1f  %sAccuracy:%s %.1f%%  %sErrors:%s %d  %sProgress:%s %d/%d\n",
		clrDim, clrReset, fmtElapsed(elapsed),
		clrDim, clrReset, grossWPM,
		clrDim, clrReset, netWPM,
		clrDim, clrReset, acc,
		clrDim, clrReset, errors,
		clrDim, clrReset, idx, len(target))

	if finished {
		fmt.Println()
		fmt.Println("✅ Finished! Press any key to exit.")
		fmt.Println()
	}
}

func fmtElapsed(d time.Duration) string {
	sec := int(d.Seconds()) % 60
	min := int(d.Minutes())
	return fmt.Sprintf("%02d:%02d", min, sec)
}

func pickRandom(items []string) string {
	if len(items) == 0 {
		return "Nothing to type. Add more sample paragraphs!"
	}
	n := big.NewInt(int64(len(items)))
	r, err := cryptoRand.Int(cryptoRand.Reader, n)
	if err != nil {
		return items[0]
	}
	// Normalize to printable width (avoid weird runes)
	s := items[int(r.Int64())]
	// Ensure valid UTF-8 (defensive)
	if !utf8.ValidString(s) {
		return strings.ToValidUTF8(s, " ")
	}
	return s
}
