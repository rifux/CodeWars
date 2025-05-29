/*
	Return the number (count) of vowels in the given string.

	We will consider a, e, i, o, u as vowels for this Kata (but not y).

	The input string will only consist of lower case letters and/or spaces.

	(Source: https://www.codewars.com/kata/54ff3102c1bad923760001f3/train/go )
*/

package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	for _, letter := range "aeiou" {
		count += strings.Count(str, string(letter))
	}
	return count
}

func main() {
	fmt.Println(GetCount("abracadabra"))
}
