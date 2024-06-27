package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

// methods with pointer receivers take either a value or a pointer as the receiver when they are called.
// Because Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.
func (receiverPtrArg *vertex) scale(f float64) {
	receiverPtrArg.X *= f
	receiverPtrArg.Y *= f
}

// functions with a pointer argument must take a pointer
func scaleFunc(regularPrtArg *vertex, f float64) {
	regularPrtArg.X *= f
	regularPrtArg.Y *= f
}

// methods with value receivers take either a value or a pointer as the receiver when they are called.
// var v Vertex
// fmt.Println(v.Abs()) // OK
// p := &v
// fmt.Println(p.Abs()) // OK
func (receiverValueArg vertex) abs() float64 {
	return math.Sqrt(receiverValueArg.X*receiverValueArg.X + receiverValueArg.Y*receiverValueArg.Y)
}

// Functions that take a value argument must take a value of that specific type
// fmt.Println(absFunc(&v)) // Compile error!
func absFunc(regularArg vertex) float64 {
	return math.Sqrt(regularArg.X*regularArg.X + regularArg.Y*regularArg.Y)
}

func main() {
	v := vertex{3, 4}
	fmt.Println(v)

	// methods with pointer receivers take either a value or a pointer as the receiver when they are called.
	v.scale(2)
	fmt.Println(v)

	// functions with a pointer argument must take a pointer
	scaleFunc(&v, 10)
	fmt.Println(v)

	p := &vertex{4, 3}
	fmt.Println(p)
	// methods with pointer receivers take either a value or a pointer as the receiver when they are called.
	p.scale(3)
	fmt.Println(p)
	// functions with a pointer argument must take a pointer
	scaleFunc(p, 8)
	fmt.Println(v, p)
	// Because Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.

	v2 := vertex{3, 4}
	fmt.Println(v2.abs())
	fmt.Println(absFunc(v2))

	p2 := &vertex{4, 3}
	fmt.Println(p2.abs())
	fmt.Println(absFunc(*p2))
}
