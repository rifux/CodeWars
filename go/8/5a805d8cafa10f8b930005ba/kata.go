package kata

import "math"

func isPerfectSq(n int) bool {
	m := int(math.Sqrt(float64(n)))
	return math.Pow(float64(m), 2) == float64(n)
}

func NearestSq(n int) (result int) {
	for modificator := 0; ; modificator++ {
		if isPerfectSq(n + modificator) {
			result = n + modificator
			break
		} else if isPerfectSq(n - modificator) {
			result = n - modificator
			break
		}
	}
	return
}
