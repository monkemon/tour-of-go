// Solution to https://go.dev/tour/methods/20

/*
Copy your Sqrt function from the earlier exercise and modify it to return an error value.

Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.
*/

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative numbers: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	before := x
	z      := 1.0
	
	
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	for i := 0; i < 100; i++ {
		z -= (z * z - x) / (2 * z)
		if delta := z - before; delta > -0.0000001 {
			return z, nil
		}
		before = z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
