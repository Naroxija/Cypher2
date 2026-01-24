package main

import "strings"

//encrypting
func encrypt_reverse(s string) string {
	return reverseAlphabet(s)
}

//decrypting
func decrypt_reverse(s string) string {
	return reverseAlphabet(s)
}

//Actual mechanism
func reverseAlphabet(s string) string {
	var result strings.Builder

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			result.WriteRune('z' - (r - 'a'))
		} else if r >= 'A' && r <= 'Z' {
			result.WriteRune('Z' - (r - 'A'))
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()

}
