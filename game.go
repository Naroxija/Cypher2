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




func initRemaining() map[rune]bool {
	m := make(map[rune]bool)
	for c := 'A'; c <= 'Z'; c++ {
		m[c] = true
	}
	return m
}

func remainingString(m map[rune]bool) string {
	var letters []string
	for c, ok := range m {
		if ok {
			letters = append(letters, string(c))
		}
	}
	sort.Strings(letters)
	return strings.Join(letters, " ")
}

func evaluate(secret, guess string, remaining map[rune]bool) string {
	result := make([]string, 5)
	used := make([]bool, 5)

	for i := 0; i < 5; i++ {
		if guess[i] == secret[i] {
			result[i] = green + string(guess[i]) + reset
			used[i] = true
		}
	}

	for i := 0; i < 5; i++ {
		if result[i] != "" {
			continue
		}
		found := false
		for j := 0; j < 5; j++ {
			if !used[j] && guess[i] == secret[j] {
				found = true
				used[j] = true
				break
			}
		}
		if found {
			result[i] = yellow + string(guess[i]) + reset
		} else {
			result[i] = white + string(guess[i]) + reset
			remaining[rune(guess[i])] = false
		}
	}

	return strings.Join(result, "")
}
