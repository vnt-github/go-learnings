package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")

	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	// Go only runs the selected case, not all the cases that follow
	switch os := runtime.GOOS; os {

	// switch cases need not be constants, and the values involved need not be integers.
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	t := time.Now()
	// Switch without a condition is the same as switch true.
	// This construct can be a clean way to write long if-then-else chains.
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
