package main

import "fmt"

// クロージャで返してあげれば
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// クロージャで受け取れて〜みたいなのできる
// まぁfunction valuesのときと一緒
func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
