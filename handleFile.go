package main

import (
	"math/rand"
	"strings"
	"time"
)

func getRandomTest() string {
	lines := strings.Split(strings.TrimSpace(testsData), "\n")
	if len(lines) == 0 {
		return ""
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strings.TrimSpace(lines[r.Intn(len(lines))])
}
