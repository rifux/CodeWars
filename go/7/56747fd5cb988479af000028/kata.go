package kata

func GetMiddle(s string) string {
	if len(s)%2 == 0 && len(s) != 0 {
		return string(s[len(s)/2-1]) + string(s[len(s)/2])
	}
	return string(s[len(s)/2])
}
