package game

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

const (
	maxAttempts = 6

	green  = "\u001B[32m"
	yellow = "\u001B[33m"
	white  = "\u001B[37m"
	reset  = "\u001B[0m"
)

func Play(scanner *bufio.Scanner, secret string) (int, bool) {
	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")

	remaining := initRemaining()
	attempts := 0

	for attempts < maxAttempts {
		fmt.Println("Enter your guess:")
		if !scanner.Scan() {
			break
		}

		guess := strings.ToUpper(strings.TrimSpace(scanner.Text()))
		if len(guess) != 5 {
			continue
		}

		attempts++
		fmt.Printf("Feedback: %s\n", evaluate(secret, guess, remaining))
		fmt.Printf("Remaining letters: %s\n", remainingString(remaining))
		fmt.Printf("Attempts remaining:  %d\n", maxAttempts-attempts)

		if guess == secret {
			return attempts, true
		}
	}

	fmt.Println(secret)
	return attempts, false
}
