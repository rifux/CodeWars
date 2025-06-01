package kata

import (
	"fmt"
	"strconv"
)

func b2i(n string) int {
	result, _ := strconv.ParseInt(n, 2, 64)
	return int(result)
}

func Int32ToIp(n uint32) string {
	uBin := strconv.FormatInt(int64(n), 2)
	for len(uBin) < 32 {
		uBin = "0" + uBin
	}
	return fmt.Sprintf("%d.%d.%d.%d", b2i(uBin[:8]), b2i(uBin[8:16]), b2i(uBin[16:24]), b2i(uBin[24:]))
}
