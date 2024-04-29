package main

import (
	"fmt"
	"math"
	"runtime"
)

// func SqrtEasy(x float64) (a float64) {
// 	current_value := 1.0
// 	new_value := 1000000000000.0
// 	for i:=0; i < 100000000; i++ {
// 		new_value = current_value - ((current_value*current_value - x) / (2*current_value))
// 		if math.Abs(new_value - current_value) < 0.001 {
// 			return new_value
// 		}
// 		current_value = new_value
// 	}
// 	return new_value
// }

func mySqrt(x float64) (a float64) {
	currentValue := 1.0
	precision := 0.000000001
	newValue := currentValue - ((currentValue*currentValue - x) / (2 * currentValue))
	for math.Abs(newValue-currentValue) > precision {
		currentValue = newValue
		newValue = currentValue - ((currentValue*currentValue - x) / (2 * currentValue))
	}
	return currentValue
}

func main() {
	q := 10.0
	ans := mySqrt(q)
	fmt.Println(ans, ans*ans, q, runtime.GOOS)
}
