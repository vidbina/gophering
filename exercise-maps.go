package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	for _, v := range strings.Fields(s) {
		_, exists := counts[v]
		if exists {
			counts[v] = counts[v] + 1
		} else {
			counts[v] = 1
		}
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
