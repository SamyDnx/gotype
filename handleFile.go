package main

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func getRandomTest(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sc := bufio.NewScanner(file)
	var candidate string
	c := 0

	for sc.Scan() {
		line := sc.Text()
		c += 1
		if r.Intn(c) == 0 {
			candidate = line
		}
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	return candidate, nil
}
