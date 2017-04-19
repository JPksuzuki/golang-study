package main

import (
	"golang.org/x/tour/pic"
	//"math"
)

func Pic(dx, dy int) [][]uint8 {
	var picture [][]uint8
	for i := 0; i < dx; i++ {
		var row []uint8
		for j := 0; j < dy; j++ {
			row = append(row, uint8((i+j)/2))
			//row = append( row, uint8(i * j) )
			//row = append( row, uint8(math.Pow(float64(i), float64(j))) )
		}
		picture = append(picture, tmp_1Dslice)
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
