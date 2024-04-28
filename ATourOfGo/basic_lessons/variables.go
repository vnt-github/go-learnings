package main

import (
	"fmt"
)

// var statement at package, since it's without initialize, type is required
var packageBoolVariable1 bool

// A var declaration can include initializers, one per variable.
// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var i, c, python, java = 10, true, false, "no!"

func main() {
	// var statement at function level, since it's without initialize, type is required
	var functionBoolVariable1 bool
	fmt.Println(packageBoolVariable1, functionBoolVariable1, i, c, python, java)
}
