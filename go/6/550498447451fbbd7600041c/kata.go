package kata

import (
	"sort"
)

func Comp(array1 []int, array2 []int) bool {
	if len(array1) == len(array2) && array1 != nil && array2 != nil {
		array1sq := make([]int, len(array1))
		for i := range array1 {
			array1sq[i] = array1[i] * array1[i]
		}
		sort.Ints(array1sq)
		sort.Ints(array2)
		for i := range array1sq {
			if array1sq[i] != array2[i] {
				return false
			}
		}
		return true
	}
	return false
}
