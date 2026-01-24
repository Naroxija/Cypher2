package main

import "strings"

//Encrypting with Caesar
func encrypt_caesar(s string) string {
	return caesarShift(s, 3)
}

//Decrypting with Caesar
func decrypt_caesar(s string) string {
	return caesarShift(s, -3)
}

func caesarShift(s string, shift int) string {
	var result strings.Builder

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			result.WriteRune('a' + (r-'a'+rune(shift)+26)%26)
		} else if r >= 'A' && r <= 'Z' {
			result.WriteRune('A' + (r-'A'+rune(shift)+26)%26)
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()

}
