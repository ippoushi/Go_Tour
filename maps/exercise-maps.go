package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	
	m := make(map[string]int)

	for _, w := range arr {
		m[w]++
	}
	return m
} 

func main() {
	wc.Test(WordCount)
}

