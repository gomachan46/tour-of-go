package main

import "fmt"

// returnの変数に名前をつけられる
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked returnっていうんだって
}

func main() {
	fmt.Println(split(17))
}
