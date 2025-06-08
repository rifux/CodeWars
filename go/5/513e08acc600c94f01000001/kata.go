package kata

func RGB(r, g, b int) string {
	table := "0123456789ABCDEF"
	result := make([]byte, 0, 6)

	for _, color := range [3]int{r, g, b} {
		if color < 0 {
			color = 0
		}
		if color > 255 {
			color = 255
		}
		result = append(result, table[color/16], table[color%16])
	}

	return string(result)
}
