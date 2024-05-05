package main

import (
	// "golang.org/x/tour/pic"
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	toRet := make([][]uint8, dy)
	for i, _ := range toRet {
		toRet[i] = make([]uint8, dx)

	}
	return toRet
}

func main() {
	fmt.Println(Pic(2, 3))

	// pic.Show(Pic)
}
