package main

import (
	"fmt"
)

const iter int = 1000
const delta float64 = 1e-15

func myAbs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Sqrt(x float64) float64 {
	z := 1.0
	count := 0
	for n := 0; n < iter; n++ {
		count++
		prev_z := z
		z = z - (z*z-x)/(2*z)
		if myAbs(z - prev_z) < delta {
			fmt.Println(count)
			return z
		}
	}
	return z
}

func main() {
	fmt.Println("\n", Sqrt(100))
}
