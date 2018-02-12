package main

import "fmt"

// ポインタは値のメモリアドレスを指す
// *int とかやるとポインタ
// &演算子はポインタを引き出す
// *演算子はポインタの指す先の変数を示す
func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the point
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
