// solution to https://go.dev/tour/moretypes/26

/*
Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers
*/

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first := 0
	second := 1
	
	fibo := func() int {
		out := first
		first = second
		second = first + out
		return out
	}
	
	return fibo
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
