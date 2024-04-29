package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i                // point to i
	fmt.Println(&i, p, *p) // read i through the pointer
	*p = 21                // set i through the pointer
	fmt.Println(i)         // see the new value of i

	p = &j                    // point to j
	*p = *p / 37              // divide j through the pointer
	fmt.Println(&i, p, &j, j) // see the new value of j
}
