package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if !isValidArgs(args) {
		printHelp()
		return
	}

	collection := args[1]
	ensureCollectionExists(collection)

	reader := bufio.NewReader(os.Stdin)
	runMenu(collection, reader)
}

/* ---------- Argument handling ---------- */

func isValidArgs(args []string) bool {
	return len(args) == 2 && args[1] != "help"
}

func printHelp() {
	fmt.Println("Usage: ./notestool [COLLECTION]")
	fmt.Println("Manage short single-line notes stored in a text file.")
}

/* ---------- Application flow ---------- */

func runMenu(collection string, reader *bufio.Reader) {
	fmt.Println("Welcome to NoteTool!")
	fmt.Println("Collection:", collection)

	for {
		printMenu()
		choice := readInput(reader)

		if handleChoice(choice, collection, reader) {
			return
		}
	}
}

func printMenu() {
	fmt.Println()
	fmt.Println("1) Display notes")
	fmt.Println("2) Add a note")
	fmt.Println("3) Remove a note")
	fmt.Println("4) Exit")
	fmt.Print("Choose an option: ")
}

func handleChoice(choice, collection string, reader *bufio.Reader) bool {
	switch choice {
	case "1":
		displayNotes(collection)
	case "2":
		addNoteFlow(collection, reader)
	case "3":
		removeNoteFlow(collection, reader)
	case "4":
		fmt.Println("Goodbye!")
		return true
	default:
		fmt.Println("Invalid option.")
	}
	return false
}

/* ---------- Input helpers ---------- */

func readInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

/* ---------- File operations ---------- */

func ensureCollectionExists(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error creating collection:", err)
		os.Exit(1)
	}
	file.Close()
}

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

func saveNotes(filename string, notes []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error saving notes:", err)
		return
	}
	defer file.Close()

	for _, note := range notes {
		fmt.Fprintln(file, note)
	}
}

func appendNote(filename, note string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error adding note:", err)
		return
	}
	defer file.Close()

	fmt.Fprintln(file, note)
}

/* ---------- Feature logic ---------- */

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

func addNoteFlow(filename string, reader *bufio.Reader) {
	fmt.Print("Enter note: ")
	note := readInput(reader)

	if note == "" {
		fmt.Println("Empty note not added.")
		return
	}

	appendNote(filename, note)
	fmt.Println("Note added.")
}

func removeNoteFlow(filename string, reader *bufio.Reader) {
	notes := loadNotes(filename)

	if len(notes) == 0 {
		fmt.Println("No notes to remove.")
		return
	}

	displayNotes(filename)
	fmt.Print("Enter note number to remove: ")

	input := readInput(reader)
	index, err := strconv.Atoi(input)

	if err != nil || index < 1 || index > len(notes) {
		fmt.Println("Invalid number.")
		return
	}

	notes = append(notes[:index-1], notes[index:]...)
	saveNotes(filename, notes)

	fmt.Println("Note removed.")
}
