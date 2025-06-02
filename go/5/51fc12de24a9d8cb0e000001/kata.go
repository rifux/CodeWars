package kata

import (
	"fmt"
	"strings"
)

func ValidISBN10(isbn string) bool {
	isbn = strings.ToUpper(isbn)

	dict := map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'X': 10,
	}

	var sum int

	for index, letter := range isbn {
		num, ok := dict[letter]
		if ok {
			sum += (index + 1) * num
		} else {
			return false
		}
	}

	if strings.Count(isbn, "X") == 1 && isbn[len(isbn)-1] != 'X' {
		return false
	}

	fmt.Println(isbn, sum, sum%11)
	return sum%11 == 0 && len(isbn) == 10 && strings.Count(isbn, "X") <= 1
}
