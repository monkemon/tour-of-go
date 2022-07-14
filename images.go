// solution to https://go.dev/tour/methods/25

package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct{
	w int
	h int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0,0, i.w, i.h)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{200,200}
	pic.ShowImage(m)
}
