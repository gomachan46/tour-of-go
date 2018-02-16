package main

import (
	"fmt"
	"math"
)

// func(float64, float64)みたいな指定ができる
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// 関数も変数なので関数の引数に取ることもできるし戻り値にすることもできるよ
func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
