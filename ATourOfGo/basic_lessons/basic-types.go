/*
Go's basic types are:

bool

string

// The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// The interface{} type (the empty interface) represents any type in Go.
// It is the interface that has no methods. All types implement zero or more interfaces.
// Go 1.18 introduces any as an alias to interface{}.
// here AnyType implements interface any which is an alias of interface{}
func printTypeAndValue[AnyType any](name string, x AnyType) {
	fmt.Printf("%s Type %T: value:|%v|\n", name, x, x)
}

// variable declarations may be "factored" into blocks, as with import statements.
var (
	ToBe bool = false
	MaxInt uint64 = 1<<64-1
	complexEx complex128 = cmplx.Sqrt(-5 + 12i)
	/*
	Variables declared without an explicit initial value are given their zero value.

	The zero value is:

	0 for numeric types,
	false for the boolean type, and
	"" (the empty string) for strings.
	*/
	i int
	f float64
	b bool
	s string

	x int = 3
	y int = 4
)



func main() {
	printTypeAndValue("ToBe", ToBe)
	printTypeAndValue("MaxInt", MaxInt)
	printTypeAndValue("complex", complexEx)
	printTypeAndValue("i", i)
	printTypeAndValue("f", f)
	printTypeAndValue("b", b)
	printTypeAndValue("s", s)
	
	// assignment between items of different type requires an explicit conversion
	f := math.Sqrt(float64(x*x + y*y))
	printTypeAndValue("f", f)

	var z uint = uint(f)
	printTypeAndValue("z", z)

	/*
	When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.
	When the right hand side of the declaration is typed, the new variable is of that same type:
	*/

	var i2 int
	j := i2 // j is an int
	printTypeAndValue("j", j)

	// But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:
	i3 := 42           // int
	f2 := 3.142        // float64
	g2 := 0.867 + 0.5i // complex128
	printTypeAndValue("i3", i3)
	printTypeAndValue("f2", f2)
	printTypeAndValue("g2", g2)

	v1 := 42 // change me!
	fmt.Printf("v1 is of type %T value: %v\n", v1, v1)
	v2 := math.Sqrt(42) // changed
	fmt.Printf("v2 is of type %T value: %v\n", v2, v2)
}