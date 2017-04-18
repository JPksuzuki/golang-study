package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	count := 0
	Sn_1 := 1
	Sn_2 := 0
	Sn := 0
	return func() int {
		if count == 0 {
			count++
			return Sn_2
		}
		if count == 1 {
			count++
			return Sn_1
		}

		count++
		Sn = Sn_1 + Sn_2
		Sn_2 = Sn_1
		Sn_1 = Sn
		return Sn
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
