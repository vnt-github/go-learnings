package main

import "fmt"

func main() {
	fmt.Println("counting")

	// A defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

	for i := 0; i <= 10; i++ {
		// Deferred function calls are pushed onto a stack. 
		// When a function returns, its deferred calls are executed in last-in-first-out order.
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
