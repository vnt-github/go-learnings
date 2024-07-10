package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// This needs specfifically a &Vertex as argument during call.
// ie. functions with a pointer argument must take a pointer
func ScaleFunc(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

// BUT
// methods with pointer receiver can take either a pointer ie. &v.Scale
// or value ie. v.Scale as the receiver/
// Since Go interprets the statement v.Scale as (&v).Scale
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// SIMILARLY, while

// Functions that take a value argument must take a value of that specific type.
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// methods with value receiver can take either a value ie. v.Abs
// or pointers ie. p := &v; p.Abs as the receiver.
// As in this case, the method call p.Abs() is interpreted as (*p).Abs().
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// NOTE: In general, all methods on a given type should have either value or
// pointer receivers, but not a mixture of both. Prefer pointer receiver.
func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	fmt.Println(v)
	ScaleFunc(&v, 10)
	fmt.Println(v)

	p := &Vertex{4, 3}
	p.Scale(3)
	fmt.Println(p)
	ScaleFunc(p, 8)
	fmt.Println(p)

	v2 := Vertex{3, 4}
	fmt.Println(v2.Abs())
	fmt.Println((*(&v2)).Abs())
	pV2 := &v2
	fmt.Println(pV2.Abs())
	fmt.Println(AbsFunc(v2))

	p2 := &Vertex{4, 3}
	fmt.Println(p2.Abs())
	fmt.Println(AbsFunc(*p2))

}
