package main

import "fmt"

// 呼び出し元の関数の終わりにdeferへ渡した関数を実行する
func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
