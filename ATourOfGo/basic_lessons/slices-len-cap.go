package main

import (
	"fmt"
)

func printSliceLengthCapacity(slice []int) {
	fmt.Printf("slice %v len=%d cap=%d is nill: %t\n", slice, len(slice), cap(slice), slice == nil)
}

func main() {
	// var a [6]int and a = [6]int{2, 3, 5, 7, 11, 13} is eq to below
	a := [6]int{2, 3, 5, 7, 11, 13}
	s := a[:]
	printSliceLengthCapacity(s)

	// NOTE here we use = and not := because that will be creating s again
	s = s[:0]
	printSliceLengthCapacity(s)

	fmt.Println(s == nil)
	if s == nil {
		fmt.Println("s is nil", s)
	}

	s = s[:4]
	printSliceLengthCapacity(s)

	s = s[2:]
	printSliceLengthCapacity(s)

	// This below will fail since 6 > cap(s) ie. 2
	// s = s[:6]
	// printSlice(s)

	s = s[2:4]

	printSliceLengthCapacity(s)

	s = s[2:]
	printSliceLengthCapacity(s)
	// A nil slice has a length and capacity of 0 and has no underlying array.

	var nillSlice []int
	printSliceLengthCapacity(nillSlice)
	if nillSlice == nil {
		fmt.Println("nil!")
	}
}
