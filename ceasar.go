package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	argumentsPassed := os.Args[1:]

	if len(argumentsPassed) != 1 || argumentsPassed[0] == "help" {
		fmt.Println("Usage: ./notestool [COLLECTION]")
		return
	}

	fileName := strings.TrimSpace(argumentsPassed[0])

	// Create collection if it does not exist
	ensureFileExists(fileName)

	for {
		operation := Prompter()
		operation = IsValidOperation(operation)

		if operation == "4" {
			fmt.Println("Exit.")
			break
		}

		ExecuteGivenOperation(operation, fileName)
	}
}

/* ---------- File existence ---------- */

func ensureFileExists(fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error creating collection:", err)
		os.Exit(1)
	}
	file.Close()
}

/* ---------- Menu ---------- */

func Prompter() string {
	fmt.Println("\nWelcome to the notes tool!")
	fmt.Println("Select operation:")
	fmt.Println("1. Show notes")
	fmt.Println("2. Add a note")
	fmt.Println("3. Delete a note")
	fmt.Println("4. Exit")
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func IsValidOperation(operation string) string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		operation = strings.TrimSpace(operation)
		n, err := strconv.Atoi(operation)

		if err == nil && n >= 1 && n <= 4 {
			return operation
		}

		fmt.Println("Please choose a number between 1 and 4:")
		scanner.Scan()
		operation = scanner.Text()
	}
}

/* ---------- Execution ---------- */

func ExecuteGivenOperation(operation string, fileName string) {
	switch operation {
	case "1":
		displayNotes(fileName)
	case "2":
		Write(fileName)
	case "3":
		removeNote(fileName)
	case "4":
		fmt.Println("Exit.")
	}
}

/* ---------- File logic ---------- */

func loadNotes(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	var notes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}
	return notes
}

func displayNotes(filename string) {
	notes := loadNotes(filename)

	if len(notes) == 0 {
		fmt.Println("No notes found.")
		return
	}

	for i, note := range notes {
		fmt.Printf("%d: %s\n", i+1, note)
	}
}

/* ---------- Writing ---------- */

func Write(fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error writing note:", err)
		return
	}
	defer file.Close()

	fmt.Print("Enter a note: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		fmt.Fprintln(file, scanner.Text())
		fmt.Println("Note added.")
	}
}

/* ---------- Removing ---------- */

func removeNote(fileName string) {
	notes := loadNotes(fileName)

	if len(notes) == 0 {
		fmt.Println("No notes to remove.")
		return
	}

	for i, note := range notes {
		fmt.Printf("%d: %s\n", i+1, note)
	}

	fmt.Print("Enter note number to delete: ")
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}

	input := strings.TrimSpace(scanner.Text())
	index, err := strconv.Atoi(input)

	if err != nil || index < 1 || index > len(notes) {
		fmt.Println("Invalid note number.")
		return
	}

	notes = append(notes[:index-1], notes[index:]...)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error updating file:", err)
		return
	}
	defer file.Close()

	for _, note := range notes {
		fmt.Fprintln(file, note)
	}

	fmt.Println("Note removed.")
}
