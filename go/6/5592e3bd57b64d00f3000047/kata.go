package kata

func FindNb(m int) int {
	for n, sum := 0, 0; sum < m; n++ {
		sum += n * n * n
		if sum == m {
			return n
		}
	}
	return -1
}
