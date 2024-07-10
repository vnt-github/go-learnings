package main

import (
	"fmt"
	"math"
)

// Interface defines a type and a set of method signatures
// value of an interface is any type that implements these methods
type Abser interface{
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f);
	}
	return float64(f);
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser;

	f := MyFloat(-math.Sqrt2)
	a = f
	fmt.Println(a.Abs())

	v := Vertex{3, 4}
	a = &v
	fmt.Println(a.Abs())

	// Bellow with give the following error
	// 'Vertex does not implement Abser (method Abs has pointer receiver)'
	// a = v
	// fmt.Println(a.Abs())
}
