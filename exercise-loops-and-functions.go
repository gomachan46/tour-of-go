package main

import (
	"fmt"
	"math"
)

// 平方根の計算実装問題
func Sqrt(x float64) float64 {
	z := 1.0
	for {
		before := z
		z -= (z*z - x) / (2 * z)
		// ごくわずかな変化しかなくなったら
		// ここもうちょっと良く出来るんかな
		if math.Abs(z-before) <= 0.000000001 {
			return z
		}
	}
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%v: %g\n", i, math.Sqrt(float64(i)))
		fmt.Printf("%v: %g\n", i, Sqrt(float64(i)))
	}
}
