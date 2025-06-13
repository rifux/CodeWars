package kata

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) []Tuple {
	res := make([]Tuple, 0, len(text))
	table := make(map[rune]int)

	for _, char := range text {
		if index, ok := table[char]; ok {
			res[index].Count++
		} else {
			table[char] = len(res)
			res = append(res, Tuple{Char: char, Count: 1})
		}
	}

	return res
}
