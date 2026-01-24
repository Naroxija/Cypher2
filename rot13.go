package main

import "strings"

//Encrypt with rot
func encrypt_rot13(s string) string {
	return rot13(s)
}

//Decrzpt with rot
func decrypt_rot13(s string) string {
	return rot13(s)
}

func rot13(s string) string {
	var result strings.Builder

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			result.WriteRune('a' + (r-'a'+13)%26)
		} else if r >= 'A' && r <= 'Z' {
			result.WriteRune('A' + (r-'A'+13)%26)
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}
