// Implementation of Map exercise from Golang's official tutorial
// https://tour.golang.org/moretypes/23

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	wordList := strings.Fields(s)
	
	for _,word := range wordList {
		wordCount[word]++
	}
	
	return wordCount
}

func main() {
	wc.Test(WordCount)
}