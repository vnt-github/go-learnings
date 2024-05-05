package main

import (
	"fmt"
	"math"
)

func printableSqrt(x float64) string {
	if x < 0 {
		return printableSqrt(-x) + "i"
	}
	ans := math.Sqrt(x)
	return string(fmt.Sprintf("%.16f", ans))
}

func myPow(x, n, lim float64) float64 {
	var ans float64
	for ans = 1; ans <= lim && n > 0; n-- {
		ans *= x
	}
	return min(ans, lim)
}

func myPow2(x, n, lim float64) float64 {
	if ans := math.Pow(x, n); ans < lim {
		return ans
	} else {
		// NOTE: ans is still avaible in the else block
		fmt.Printf("%g >= %g\n", ans, lim)
	}
	return lim
}

func main() {
	var floatArr [10]float64
	for i := 1; i < 10; i++ {
		floatArr[i] = float64(i * i)
	}

	fmt.Println(floatArr)

	for index, each := range floatArr {
		fmt.Println(index+1, "sqrt of ", each, ":", printableSqrt(each))
	}

	fmt.Println("myPow(3, 2, 10)", myPow(3, 2, 10), myPow2(3, 2, 10))
	fmt.Println("myPow(3, 3, 10)", myPow(3, 3, 10), myPow2(3, 3, 10))
	fmt.Println("myPow(3, 3, 20)", myPow(3, 3, 20), myPow2(3, 3, 20))
	fmt.Println("myPow(3, 3, 30)", myPow(3, 3, 30), myPow2(3, 3, 30))
}
