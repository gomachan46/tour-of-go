package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	r := make([][]uint8, dy)
	for i := range r {
		x := make([]uint8, dx)
		for j := range x {
			x[j] = uint8(i * j)
		}
		r[i] = x
	}

	return r
}

func main() {
	pic.Show(Pic)
}
