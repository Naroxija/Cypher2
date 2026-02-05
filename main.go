package main

import (
	"bufio"
	"os"
	"strconv"

	"koodWordle/game"
	"koodWordle/io"
	"koodWordle/model"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	words := io.LoadWords("wordle-words.txt")
	if len(words) == 0 || index < 0 || index >= len(words) {
		return
	}

	secret := words[index]
	scanner := bufio.NewScanner(os.Stdin)

	username := io.ReadUsername(scanner)
	if username == "" {
		return
	}

	user := model.NewUser(username)

	attempts, won := game.Play(scanner, secret)

	io.WriteStats(user.Username, secret, attempts, won)

	if io.AskShowStats(scanner) {
		io.ShowStats(user.Username)
	}

	io.WaitForExit(scanner)
}
