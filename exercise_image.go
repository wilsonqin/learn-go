package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	height int
	width int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0,0,i.width,i.height)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	v := uint8( y^2 + 2*x*y + x^2 )
	return color.RGBA{v, v, 255, 255}
}

func RunImage() {
	m := Image{50,50}
	pic.ShowImage(m)
}
