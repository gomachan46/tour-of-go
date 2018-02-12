package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42                       // create
	fmt.Println("The value:", m["Answer"]) // read

	m["Answer"] = 48 // update
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") // delete
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // okで要素の有無判定が可能
	fmt.Println("The value:", v, "Present?", ok)
}
