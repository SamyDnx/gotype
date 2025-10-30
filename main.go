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

// ansi code for colors and reset
const Red string = "\033[31m"
const Green string = "\033[32m"
const Neural string = "\033[37m"
const Reset string = "\033[2K"

func main() {
	text := []rune("The quick brown fox jumps over the lazy dog")
	user := UserTextProgress{
		progress: [256]int8{},
		idx:      0,
	}

	nonLetterRune := 0
	timeStarted := false
	var start time.Time

	oldState, err := enableRawMode()
	if err != nil {
		fmt.Printf("Failed to enable raw mode: %v\n", err)
		return
	}
	defer disableRawMode(oldState)

	// 1st text print
	for _, r := range text {
		if (r < 65 && r > 90) || (r < 97 && r > 122) {
			nonLetterRune += 1
		}
		fmt.Printf("%s", string(r))
	}
	for {
		if user.idx >= len(text) {
			end := time.Now()
			delta := end.Sub(start)
			disableRawMode(oldState)
			fmt.Printf("\nTest finished in %.2f seconds\n ", delta.Seconds())
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
				fmt.Println("Exiting...")
				return
			}
		}
		renderText(&user, text)
		user.idx += 1
	}
}

// render text with according colors
func renderText(user *UserTextProgress, text []rune) {
	// clear line and put the cursor at the start of line
	fmt.Printf("%s\r", Reset)
	for i, r := range text {
		if user.progress[i] > 0 {
			fmt.Printf("%s%s%s", Green, string(r), Neural)
		} else if user.progress[i] < 0 {
			fmt.Printf("%s%s%s", Red, string(r), Neural)
		} else {
			fmt.Printf("%s", string(r))
		}
	}
}
