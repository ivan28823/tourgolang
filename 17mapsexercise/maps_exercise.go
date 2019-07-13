package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	p := make(map[string]int)

	str := strings.Fields(s)

	for _, v := range str {
		p[v]++
	}

	return p
}

func main() {
	wc.Test(WordCount)
}
