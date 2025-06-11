package kata

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) (result PosPeaks) {
	if len(array) < 3 {
		return
	}

	var prevIndex int
	for i := 0; i < len(array)-1; i++ {
		if array[i] != array[i+1] {
			if array[prevIndex] < array[i] && array[i] > array[i+1] {
				result.Pos = append(result.Pos, prevIndex+1)
				result.Peaks = append(result.Peaks, array[prevIndex+1])
			}
			prevIndex = i
		}
	}

	return
}
