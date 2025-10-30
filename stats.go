package main

import "fmt"

// ascii codes
const SPACE rune = 32
const UPPER_A rune = 65
const UPPER_Z rune = 95
const LOWER_A rune = 97
const LOWER_Z rune = 122

// getLettersCount returns the number of letters and the numbers of letters correctly typed
func getLetterCount(user *UserTextProgress, text []rune) (int, int) {
	letterCount := 0
	correctLetters := 0
	for i, l := range text {
		if (l >= UPPER_A && l <= UPPER_Z) || (l >= LOWER_A && l <= LOWER_Z) {
			letterCount += 1
			if user.progress[i] == 1 {
				correctLetters += 1
			}
		}
	}
	return letterCount, correctLetters
}

func getWordCount(text []rune) int {
	wordCount := 1
	for _, r := range text {
		// check for spaces
		if r == SPACE {
			wordCount += 1
		}
	}
	return wordCount
}

// getWPM returns the WPM and the number of words correctly typed
func getWPM(user *UserTextProgress, text []rune, delta float64) (float64, int) {
	correctWords := 0
	wordHasError := false

	for i, p := range user.progress {
		if i >= len(text) {
			break
		}

		char := text[i]
		if char != SPACE {
			if p == -1 {
				wordHasError = true
			}
		}
		if char == SPACE {
			if !wordHasError {
				correctWords += 1
			}
			wordHasError = false
		}
	}
	if !wordHasError {
		correctWords++
	}

	wpm := (float64(correctWords) / delta) * 60
	return wpm, correctWords
}

// get Accuracy returns the % of words correctly typed and the % of letters correctly typed
func getAccuracy(user *UserTextProgress, text []rune, correctWords int) (float64, float64) {
	wordCount := getWordCount(text)
	wordPercent := (float64(correctWords) / float64(wordCount)) * 100

	letterCount, correctLetters := getLetterCount(user, text)
	letterPercent := (float64(correctLetters) / float64(letterCount)) * 100

	return wordPercent, letterPercent
}

func printStats(user *UserTextProgress, text []rune, delta float64) {
	wpm, correctWords := getWPM(user, text, delta)
	wordCount := getWordCount(text)
	wordPercent, letterPercent := getAccuracy(user, text, correctWords)
	letterCount, correctLetter := getLetterCount(user, text)

	fmt.Printf("WPM: %.2f\n", wpm)
	fmt.Printf("Word accuracy: %.2f%% --> %d/%d\n", wordPercent, correctWords, wordCount)
	fmt.Printf("Letter accuracy: %.2f%% --> %d/%d\n", letterPercent, correctLetter, letterCount)
}
