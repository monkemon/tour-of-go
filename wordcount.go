package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	out := make(map[string]int)
	
	for _, word := range(strings.Fields(s)){
		if _, ok := out[word]; ok == true {
			out[word] ++
		} else {
			out[word] = 1
		}
	}
	return out
}

func main() {
	wc.Test(WordCount)
}
