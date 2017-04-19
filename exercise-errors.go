package main

import (
	"fmt"
	"math"
)

const iter int = 1000
const delta float64 = 1e-15

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return -1, ErrNegativeSqrt(x) //return -1 as error
	}
	z := 1.0
	count := 0
	for n := 0; n < iter; n++ {
		count++
		prev_z := z
		z = z - (z*z-x)/(2*z)
		if math.Abs(z-prev_z) < delta {
			fmt.Println(count)
			return z, nil
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
