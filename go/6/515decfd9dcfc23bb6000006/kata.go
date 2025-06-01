package kata

import (
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
	splits := strings.Split(ip, ".")
	if len(splits) == 4 {
		for _, number := range splits {
			numberInt, err := strconv.Atoi(number)
			if err != nil || numberInt > 255 || numberInt < 0 || len(number) != len(strconv.Itoa(numberInt)) {
				return false
			}
		}
		return true
	}
	return false
}
