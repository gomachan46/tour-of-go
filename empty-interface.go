package main

import "fmt"

// 空のインタフェースは任意の型の値を保持できる(すべての型はゼロ個のメソッドを実装していると言える)
// 空のインタフェースは未知の型の値を扱うコードで使用する。fmt.Printとか
func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
