package main

import "fmt"

// nilスライスは0の長さと容量をもち、元となる配列を持っていない
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
