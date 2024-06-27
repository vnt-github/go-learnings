// There are two reasons to use a pointer receiver.
// 1. method can modify the value that its receiver points to.
// 2. avoid copying the value on each method call. This can be more efficient if the receiver is a large struct,
// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

func (v *vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *vertex) scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	a := vertex{3, 4}
	fmt.Println(a, a.abs())
	a.scale(10)
	fmt.Println(a, a.abs())

	v := &vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.abs())
	v.scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.abs())
}
