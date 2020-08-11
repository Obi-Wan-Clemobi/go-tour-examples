package main

import (
	_"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := make(map[string]int)
	for _, field := range strings.Fields(s) {
		fields[field] += 1
	}
	return fields
}

func main() {
	wc.Test(WordCount)
}
