package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(
	singleLineWords string,
) map[string]int {
	tokens := strings.Fields(singleLineWords)
	wordCounts := make(map[string]int)
	
	for _, token := range tokens {
		wordCounts[token] += 1
	}
	
	return wordCounts
}

func main() {
	wc.Test(WordCount)
	
	WordCount("")
}

