package kata

func HighestRank(nums []int) int {
	table := make(map[int]int)

	for _, number := range nums {
		table[number]++
	}

	var maxInt, maxCnt int
	for key, value := range table {
		if maxCnt < value || maxCnt == value && maxInt < key {
			maxInt = key
			maxCnt = value
		}
	}

	return maxInt
}
