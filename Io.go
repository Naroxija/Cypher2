package io

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadUsername(scanner *bufio.Scanner) string {
	fmt.Println("Enter your username:")
	if !scanner.Scan() {
		return ""
	}
	return strings.TrimSpace(scanner.Text())
}

func AskShowStats(scanner *bufio.Scanner) bool {
	fmt.Println("Do you want to see your stats? (yes/no):")
	if !scanner.Scan() {
		return false
	}
	return strings.ToLower(strings.TrimSpace(scanner.Text())) == "yes"
}

func WaitForExit(scanner *bufio.Scanner) {
	fmt.Println("Press Enter to exit...")
	scanner.Scan()
}

func LoadWords(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if len(word) == 5 {
			words = append(words, strings.ToUpper(word))
		}
	}
	return words
}

func WriteStats(username, secret string, attempts int, won bool) {
	file, _ := os.OpenFile("stats.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	result := "loss"
	if won {
		result = "win"
	}

	writer.Write([]string{
		username,
		secret,
		strconv.Itoa(attempts),
		result,
	})
}

func ShowStats(username string) {
	file, err := os.Open("stats.csv")
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, _ := reader.ReadAll()

	games, wins, totalAttempts := 0, 0, 0

	for _, row := range rows {
		if row[0] != username {
			continue
		}
		games++
		a, _ := strconv.Atoi(row[2])
		totalAttempts += a
		if row[3] == "win" {
			wins++
		}
	}

	avg := 0.0
	if games > 0 {
		avg = float64(totalAttempts) / float64(games)
	}

	fmt.Printf("Stats for %s:\n", username)
	fmt.Printf("Games played: %d\n", games)
	fmt.Printf("Games won: %d\n", wins)
	fmt.Printf("Average attempts per game: %.2f\n", avg)
}
