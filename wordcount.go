// solution to go wordcount
// https://go.dev/tour/moretypes/23

/*
Implement WordCount. It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.
*/

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
