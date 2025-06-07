package kata

import (
	"strings"
	"unicode"
)

func Order(sentence string) string {
	words := strings.Split(sentence, " ")
	ordered := make([]string, len(words))

	for _, word := range words {
		for _, letter := range word {
			if unicode.IsDigit(letter) {
				ordered[int(byte(letter)-'0')-1] = word
				continue
			}
		}
	}

	return strings.Join(ordered, " ")
}
