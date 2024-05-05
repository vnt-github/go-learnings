package main

import (
	"fmt"
	"strings"
)

func printxob(xob [][]string) {
	fmt.Println()
	for i := 0; i < len(xob); i++ {
		fmt.Printf("%s\n", strings.Join(xob[i], " | "))
		if i < len(xob)-1 {
			fmt.Printf("----------\n")
		}
	}
	fmt.Println()
}

func main() {
	xob := [][]string{
		{"X", " ", "X"},
		{" ", "O", " "},
		{" ", "O", "X"},
	}
	fmt.Println(xob)
	printxob(xob)
}
