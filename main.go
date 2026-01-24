package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Cypher tool!")

	toEncrypt, encoding, message := getInput()

	var result string

	switch encoding {

	case "rot13":
		if toEncrypt {
			result = encrypt_rot13(message)
		} else {
			result = decrypt_rot13(message)
		}
	case "reverse":
		if toEncrypt {
			result = encrypt_reverse(message)
		} else {
			result = decrypt_reverse(message)
		}
	case "caesar":
		if toEncrypt {
			result = encrypt_caesar(message)
		} else {
			result = decrypt_caesar(message)
		}
	}

	fmt.Println("\n Result:")
	fmt.Println(result)
}
