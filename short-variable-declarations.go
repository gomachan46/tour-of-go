package main

import "fmt"

// :=でvarの代わりに型宣言できる
// 関数外ではだめ
func main() {
	var i, j = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
