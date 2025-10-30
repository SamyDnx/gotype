package main

import (
	"fmt"
	"time"
)

// store user progress
type UserTextProgress struct {
	progress [256]int8
	idx      int
}

// ansi code for colors and RESET
const RED string = "\033[31m"
const GREEN string = "\033[32m"
const NEUTRAL string = "\033[37m"
const RESET string = "\033[2K"
const CLEAR string = "\033[2J"

// path to file with tests sentences
const PATH string = "tests.txt"

func main() {
	test, err := getRandomTest(PATH)
	if err != nil {
		fmt.Printf("Failed reading the file: %v\n", err)
	}

	text := []rune(test)
	user := UserTextProgress{
		progress: [256]int8{},
		idx:      0,
	}

	timeStarted := false
	var start time.Time

	oldState, err := enableRawMode()
	if err != nil {
		fmt.Printf("Failed to enable raw mode: %v\n", err)
		return
	}
	defer disableRawMode(oldState)

	// clear window
	fmt.Print(CLEAR)
	// 1st text print
	for i, r := range text {
		fmt.Printf("\033[1;%dH%s", i+1, string(r))
	}
	// set cursor at start of terminal
	fmt.Print("\033[1;1H")

	for {
		if user.idx >= len(text) {
			end := time.Now()
			delta := end.Sub(start)
			disableRawMode(oldState)
			fmt.Printf("\nTest finished in %.2f seconds\n", delta.Seconds())
			printStats(&user, text, delta.Seconds())
			return
		}

		buf, n := readInput()

		// check timer
		if !timeStarted {
			start = time.Now()
			timeStarted = true
		}

		// check typed letter
		if n > 0 {
			if buf[0] == byte(text[user.idx]) {
				user.progress[user.idx] = 1
			} else {
				user.progress[user.idx] = -1
			}

			// extiting with CTRL+D (4)
			if buf[0] == 4 {
				fmt.Println("\nExiting...")
				return
			}
		}
		renderText(&user, text)
		user.idx += 1
	}
}

// render text with according colors
func renderText(user *UserTextProgress, text []rune) {
	// clear widow and put the cursor at the start of terminal
	fmt.Printf("%s\033[1;1H", CLEAR)
	for i, r := range text {
		if user.progress[i] > 0 {
			fmt.Printf("%s%s%s", GREEN, string(r), NEUTRAL)
		} else if user.progress[i] < 0 {
			fmt.Printf("%s%s%s", RED, string(r), NEUTRAL)
		} else {
			fmt.Printf("%s", string(r))
		}
	}
	fmt.Printf("\033[1;%dH", user.idx+2)
}
