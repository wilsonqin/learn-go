package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string] int)
	words := strings.Fields(s)
	for _, word := range words {
		m[word] = m[word] + 1
	}
	return m
}

func RunWordCount() {
	wc.Test(WordCount)
}
