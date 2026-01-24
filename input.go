package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// We are getting the input with this function
func getInput() (toEncrypt bool, encoding string, message string) {
	reader := bufio.NewReader(os.Stdin)

	// Input for the encryption/decryption choice
	for {
		fmt.Print("Choose operation (encrypt/decrypt): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "encrypt" {
			toEncrypt = true
			break
		} else if input == "decrypt" {
			toEncrypt = false
			break
		}
	}

	// Input for the different optionz
	for {
		fmt.Print("Choose cipher (rot13/reverse/caesar): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "rot13" || input == "reverse" || input == "caesar" {
			encoding = input
			break
		}
	}

	//Return
	for {
		fmt.Print("Enter message: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input != "" {
			message = input
			break
		}
	}

	return
}
