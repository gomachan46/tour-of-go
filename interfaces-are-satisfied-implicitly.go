package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// TにM()を実装してるから
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	// インタフェースを満たしてる
	var i I = T{"Hello"}
	i.M()
}
