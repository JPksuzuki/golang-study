package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var want = make(map[string]int)

	var words []string = strings.Fields(s)

	for _, word := range words {
		_, isExist := want[word]
		if !isExist {
			want[word] = 1
		} else {
			want[word]++
		}
	}
	return want
}

func main() {
	wc.Test(WordCount)
}
