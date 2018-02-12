package main

import "fmt"

// 全部同じ意味
// s[0:6]
// s[0:]
// s[:6]
// s[:]
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
