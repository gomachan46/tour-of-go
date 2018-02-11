package main

import "fmt"

// func 関数名(引数名 型) 返り値の型
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
