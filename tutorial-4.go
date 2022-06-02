package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)

	words := strings.Fields(s)

	for i := range words {

		var word string
		var currentCount int
		word = words[i]
		currentCount = result[word]
		result[word] = currentCount + 1

	}

	return result
}

func main() {
	wc.Test(WordCount)
}
