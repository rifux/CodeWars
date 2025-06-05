package kata

func FindUniq(arr []float32) float32 {
	counter := make(map[float32]int)
	for _, num := range arr {
		counter[num]++
	}
	for num, count := range counter {
		if count == 1 {
			return num
		}
	}
	return 0
}
