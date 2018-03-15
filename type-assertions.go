package main

import "fmt"

// 型アサーション
// t := i.(T)
// この文は、インタフェースの値iが具体的な型Tを保持し、元になるTの値を変数tに代入するという意味
// インタフェースの値が特定の型を保持しているかどうかをテストするために型アサーションは2つの値(元になる値とアサーションの結果)を返すことが出来る
// t, ok := i.(T)
// iがTを保持していればtは元になる値、okはtrue
// okまで結果を取るようにすればpanicは起きない
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
