package main

import "golang.org/x/tour/pic"

// Pic todo
func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		ret[i] = make([]uint8, dx)
	}

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			ret[x][y] = uint8((x*x + y*y))
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
