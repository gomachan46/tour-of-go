package main

import "fmt"

// 初期化も自動で出来るよ
// 初期化したときは型いらないよ(値から勝手に良い感じにしてくれるよ)
var i, j = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
