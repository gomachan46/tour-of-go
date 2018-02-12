package main

import "fmt"

func main() {
	sum := 1
	// セミコロンも省略可能
	// goではwhileもforで表現
	for sum < 1000 {
		sum += sum
	}
	fmt.Println((sum))
}
