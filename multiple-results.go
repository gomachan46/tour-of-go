package main

import "fmt"

// 複数返せる
// 型の指定は引数で同じ型なら最後にまとめてかける
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
