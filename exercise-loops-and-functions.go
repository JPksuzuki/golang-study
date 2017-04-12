package main

import (
	"fmt"
)

var iter int = 1000
var delta float64 = 0.00000000001

func myAbs(x float64) float64 {
	if x < 0 { return -x }
	return x
}

func Sqrt(x float64) float64 {
	z := 1.0
	count := 0
	for n := 0; n < iter; n++ {
		count++		
		z_tmp := z
		z = z - (z * z - x)/(2 * z)
		if myAbs(z - z_tmp) < delta {
			fmt.Println(count)
			return z
		} 
	}
	return z
}

func main() {
	fmt.Println("\n",Sqrt(100))
}
