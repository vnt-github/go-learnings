package main

import (
	"fmt"
)

func main() {
	// An array's length is part of its type, so arrays cannot be resized
	var greetingWords [2]string

	greetingWords[0] = "Hello World"
	greetingWords[1] = "go"

	fmt.Println(greetingWords, greetingWords[0], greetingWords[1])

	// NOTE: Using {} this is similar to the struct syntax v3 = Vertex{}
	// NOTE: assignment using := is required when using {} to initialize and you can't use var, const
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// so below won't work and give error and if you do var new_primes = [6]int{2, 3, 5, 7, 11, 13} then it auto corrects to above syntax
	// var fauly_primes [6]int{2, 3, 5, 7, 11, 13}

	// this below works
	var newPrimes [6]int = [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(newPrimes)

	primesSliceIncludedStartIndex := 1
	primesSliceExcludedEndIndex := 3

	// The type []T is a slice with elements of type T.
	// A slice is formed by specifying two indices, a low and high bound
	// array[high:low] for half-open range.
	var primeSlice []int = primes[primesSliceIncludedStartIndex:primesSliceExcludedEndIndex]

	fmt.Println(primeSlice)

	// Slices are like references to arrays: Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.

	var assignedNames [4]string
	assignedNames[0] = "kriss"
	assignedNames[1] = "jorge"
	assignedNames[2] = "christina"
	assignedNames[3] = "isa"
	fmt.Println(assignedNames)

	initiatedNames := [4]string{
		"kriss",
		"empty",
		"christina",
		"isa",
	}

	fmt.Println(initiatedNames)

	// NOTE: slice has to be initialized with it's corresponding array
	// sliced_names1 []string;
	// slicedNames1 = initiatedNames[0:2]

	slicedNames1 := initiatedNames[0:2]
	slicedNames2 := initiatedNames[1:3]

	slicedNames1[1] = "Vineet"

	fmt.Println(initiatedNames, slicedNames1, slicedNames2)
}
