package kata

func WordsToMarks(s string) (sum int) {
	for _, char := range s {
		sum += (int(char) - 96)
	}
	return
}
