package main

import (
	"fmt"
)

// var statement at package, since it's without initialize, type is required
var packageBoolVariable1 bool

// A var declaration can include initializers, one per variable.
// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
var i, c, python, java = 10, true, false, "no!"

func main() {
	// var statement at function level, since it's without initialize, type is required
	var functionBoolVariable1 bool
	fmt.Println(packageBoolVariable1, functionBoolVariable1, i, c, python, java)

	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	k := 3
	fmt.Printf("Type and value can be set using ':='\nEx: for K, Type: %T Value: %v\n", k, k)

}
