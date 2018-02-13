package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	out := make(map[string]int)
	for _, w := range strings.Fields(s) {
		out[w]++
	}
	return out
}

func main() {
	wc.Test(WordCount)
}
