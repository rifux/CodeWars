package kata

import "strings"

func duplicate_count(s1 string) (counter int) {
	dict := make(map[rune]int)
	for _, char := range strings.ToLower(s1) {
		dict[char] += 1
	}
	for _, value := range dict {
		if value > 1 {
			counter++
		}
	}
	return
}
