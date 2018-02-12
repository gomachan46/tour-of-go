package main

import "fmt"

type Vertex5 struct {
	Lat, Long float64
}

var m map[string]Vertex5

// mapのゼロ値はnil
// nilマップはキーを持ってないしキーを追加することも出来ない
// 使えるようにするにはmake関数で初期化してあげる
func main() {
	m = make(map[string]Vertex5)
	m["Bell Labs"] = Vertex5{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
