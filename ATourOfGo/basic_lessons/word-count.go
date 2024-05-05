package main

import (
	"fmt"
	"strings"
	// "golang.org/x/tour/wc"
)

func wordCount(s string) map[string]int {
	// var toRet = map[string]int{}
	var toRet map[string]int
	toRet = make(map[string]int)
	// toRet := make(map[string]int)

	for _, word := range strings.Fields(s) {
		toRet[word]++
	}

	return toRet
}

func main() {
	// wc.Test(wordCount)
	fmt.Println(wordCount("a quick brown fox jumps over a lazy dog"))
}
