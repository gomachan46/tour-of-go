package main

import "fmt"

// いらないものが返ってくるときは_で捨てることができる
// もしインデックスだけ必要なら後のreturnは記述しなければok
func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
	for i := range pow {
		fmt.Printf("%d\n", i)
	}
}
