package main

import "fmt"

// スライスは配列への参照のようなもの
// スライスはデータを格納しておらず、単に元の配列の部分列を指し示しているだけ
// スライスの要素を変更すると元の配列の要素も変更される
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
