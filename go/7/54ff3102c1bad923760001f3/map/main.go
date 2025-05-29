package main

/*
Return the number (count) of vowels in the given string.

We will consider a, e, i, o, u as vowels for this Kata (but not y).

The input string will only consist of lower case letters and/or spaces.
*/

type empty struct{}

var vowels map[string]empty = map[string]empty{
	"a": {},
	"e": {},
	"i": {},
	"o": {},
	"u": {},
}

func GetCount(str string) (count int) {
	// Enter solution here
	for _, letter := range str {
		_, isVowel := vowels[string(letter)]
		if isVowel {
			count++
		}
	}
	return
}
