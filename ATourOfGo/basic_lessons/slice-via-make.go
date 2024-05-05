package main

import "fmt"

func printSliceLengthCapacity(slice []int) {
	fmt.Printf("slice: %v %d %d\n", slice, len(slice), cap(slice))
}

func main() {
	a1 := []int{0, 0, 0, 0}
	printSliceLengthCapacity(a1)

	// We can create a slice using the above syntax but what if you have to create a slice of size 100
	// It's not practicle to create {0,0,0...0} 100 times hence the below make syntax is used
	a1_better := make([]int, 100, 100)
	printSliceLengthCapacity(a1_better)

	a := make([]int, 4, 9)

	printSliceLengthCapacity(a)

	a = a[:9]
	printSliceLengthCapacity(a)

	// NOTE: since the capacity is 9 the below slice beyond 10 won't work
	// a = a[:10]

	// func append(s []T, vs ...T) []T
	// The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
	// The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
	// If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

	a = append(a, -1, -2, -3, -4)
	printSliceLengthCapacity(a)

	b := make([]int, 0, 5)

	b = []int{1, 2, 3, 4, 5}

	printSliceLengthCapacity(b)

	c := b[:2]
	printSliceLengthCapacity(c)

	d := c[2:5]
	printSliceLengthCapacity(d)
}
