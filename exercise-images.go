package main

import(
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (img Image) At(i, j int) color.Color {
	v := uint8((i + j) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
