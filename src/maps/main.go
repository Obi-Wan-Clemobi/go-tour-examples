package main

import (
	_ "fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	fields := make(map[string]int)
	for _, field := range strings.Fields(s) {
		fields[field]++
	}
	return fields
}

func main() {
	wc.Test(WordCount)
}
