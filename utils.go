package bt

import (
	"math/rand"
)

// 乱序一个整数数组
func reArrangeIntSlice(intSlc []int) {
	for i := 1; i < len(intSlc); i++ {
		r := rand.Intn(i)
		intSlc[i], intSlc[r] = intSlc[r], intSlc[i]
	}
}
