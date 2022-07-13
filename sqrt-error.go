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

func Sqrt(x float64) (float64, ErrNegativeSqrt) {
	before := x
	z      := x
	
	err := nil
	
	
	if x < 0 {
		err = ErrNegativeSqrt(x)
		return x, err
	}
	for i := 0; i < 100; i++ {
		z -= (z * z - x) / (2 * z)
		if delta := z - before; delta > -0.0000001 {
			return z, err
		}
		before = z
	}
	return z, err
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
