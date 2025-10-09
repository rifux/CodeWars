package kata

func ArrayDiff(a, b []int) []int {
	bSet := make(map[int]struct{})
	for _, item := range b {
		bSet[item] = struct{}{}
	}

	aNew := []int{}
	for _, item := range a {
		_, ok := bSet[item]
		if !ok {
			aNew = append(aNew, item)
		}
	}

	return aNew
}
