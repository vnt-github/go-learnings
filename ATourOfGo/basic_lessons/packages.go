// main package for learning about packages in Go
package main

// Every Go program is made up of packages.
// Programs start running in package main.
// This program is using the packages with import paths "fmt" and "math/rand".

// This below is factorized import
// By convention, the package name is the same as the last element of the import path.
// For instance, the "math/rand" package comprises files that begin with the statement `package rand`.
import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	// name beginning with capital letter is exported hence .Pi should be used instead of .pi
	// fmt.println(math.pi)
	fmt.Println(math.Pi)
}
