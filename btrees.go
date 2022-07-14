// solution to https://go.dev/tour/concurrency/8

package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if left := t.Left; left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if right := t.Right; right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	// following 'close()' calls might be defered in the Walk funciton
	go Walk(t1, ch1)
	
	go Walk(t2, ch2)
	
	for {
		v1, ok := <- ch1
		v2, okk:= <- ch2
		
		if v1 != v2  || ok != okk{
			return false
		}
		
		if !ok {
			break
		}
		
		
	}
	return true
	
}

func main() {
	tre := tree.New(1)
	ch := make(chan int)
	
	go Walk(tre, ch)
	
	for i := 0; i < 10; i++ {
		val, ok := <- ch
		if ok == true {
			fmt.Println("ok is:", ok, "value: ", val)
		} else {
			return
		}
	}
	
	tre = tree.New(12)
	treee := tree.New(12)
	
	fmt.Println(Same(tre, treee))
			
}
