package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

// Go does not have classes. However, you can define methods on type
// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// In this example, the Abs method has a receiver of type Vertex named v.
// This value receiver method operates on a copy of the original Vertex value
func (receverArgV vertex) Abs() float64 {
	return math.Sqrt(receverArgV.X*receverArgV.X + receverArgV.Y*receverArgV.Y)
}

// You can declare methods with pointer receivers.
// This means the receiver type has the literal syntax *T for some type T.
// (Also, T cannot itself be a pointer such as *int.)
// For example, the Scale method here is defined on *Vertex.
// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
// This Scale method must have a pointer receiver to change the Vertex value declared in the main function.
func (receiverPtrArg *vertex) Scale(f float64) {
	receiverPtrArg.X *= f
	receiverPtrArg.Y *= f
}

// Remember a method is just a function with a receiver argument.
// Here's Abs written as a regular function with no change in functionality.
func abs2(regularArg vertex) float64 {
	return math.Sqrt(regularArg.X*regularArg.X + regularArg.Y*regularArg.Y)
}

// You can declare a method on non-struct types, too.
type myFloat float64

// You can only declare a method with a receiver whose type is defined in the same package.
// You cannot declare a method with a receiver whose type is defined in another package
// (which includes the built-in types such as int).
func (f myFloat) floatAbs() myFloat {
	if f < 0 {
		return -f
	}
	return f
}

func main() {
	v := vertex{3, 4}
	fmt.Println(v.X, v.Y, v.Abs())
	fmt.Println(abs2(v))
	f := myFloat(-math.Sqrt2)
	fmt.Println(f)
	fmt.Println(f.floatAbs())
	fmt.Println(f)
	v.Scale(10)
	fmt.Println(v, v.Abs())
}
