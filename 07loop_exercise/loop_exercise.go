package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(x / 8)
	for i := 0; i < 20; i++ {
		fmt.Println(i, z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	var num float64 = 65465464.645654
	fmt.Println(Sqrt(num), math.Sqrt(num))
}
