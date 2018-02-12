package main

import "fmt"

type Vertex3 struct {
	X int
	Y int
}

// structのポインタを通じてアクセスすることも可能
func main() {
	v := Vertex3{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
