package main

import "fmt"

// 配列は固定長
// スライスは可変長で、より柔軟な配列
// []TはTのスライスという意味([6]Tとかだと配列)
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
