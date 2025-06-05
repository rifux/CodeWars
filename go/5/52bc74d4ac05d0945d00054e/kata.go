package kata

import (
	"strings"
)

func FirstNonRepeating(str string) string {
	occur := make(map[rune]int)
	for _, char := range str {
		charLow := []rune(strings.ToLower(string(char)))[0]
		occur[charLow]++
	}
	for _, char := range str {
		charLow := []rune(strings.ToLower(string(char)))[0]
		if occur[charLow] == 1 {
			return string(char)
		}
	}
	return ""
}
