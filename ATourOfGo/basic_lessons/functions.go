// basic package for learning about functions
package main

import (
	"fmt"
	"math/rand"
)

// https://go.dev/blog/gos-declaration-syntax: having type after expression makes syntax more readable as compared to C.
// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func add(x, y int) int {
	return x + y
}

// A function can return any number of results.
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// const variables of type int, const values must be defined along with declaration
	const MaxRandomIntRange int = 100

	// variables a and b for holding random values to print
	var a, b int
	a = rand.Intn(MaxRandomIntRange)
	b = rand.Intn(MaxRandomIntRange)
	fmt.Println("a:", a, "b:", b, "\nsum", add(a, b))

	// := operator is a shorthand for declaration and assignment, it automatically infers the type of the variable based on the value assigned to it.
	x, y := swap("alpha", "beta")
	fmt.Println("after swap, with := ie. declaration and assignment\nx:", x, "y:", y)

	// := is only used for declaration and assignment, for reassignment use =
}
