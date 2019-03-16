package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image the image struct
type Image struct {
	width  int
	height int
	pixels [][]color.Color
}

// NewImage the constructor of Image
func NewImage(width, height int, pixelGenerator func(x, y uint8) color.Color) *Image {
	img := &Image{width: width, height: height}
	img.pixels = make([][]color.Color, width)
	for i := 0; i < width; i++ {
		img.pixels[i] = make([]color.Color, height)
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.pixels[x][y] = pixelGenerator(uint8(x), uint8(y))
		}
	}
	return img
}

// ColorModel todo
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds todo
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

// At todo
func (img Image) At(x, y int) color.Color {
	return img.pixels[x][y]
}

func generator1(x, y uint8) color.Color {
	return color.RGBA{x, y, x + y, 255}
}

func generator2(x, y uint8) color.Color {
	return color.RGBA{x, (x + y) / 2, (x - y) / 2, 200}
}

func main() {
	m := NewImage(200, 200, generator1)
	pic.ShowImage(m)
}
