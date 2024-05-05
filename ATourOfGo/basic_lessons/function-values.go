package main

import (
	"fmt"
	"math"
)

// Functions are values too. They can be passed around just like other values.
// Function values may be used as function arguments and return values.
func computefnOn3and4(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	fmt.Println(computefnOn3and4(math.Pow))

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(computefnOn3and4(hypot))
}
