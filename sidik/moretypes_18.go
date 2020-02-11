package sidik

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	sdx := make([][]uint8, dx)
	for y := 0; y < dy; y++ {
		sdy := make([]uint8, dy)
		for x := 0; x < dx; x++ {
			sdy[x] = uint8((x+y)*4)
		}
		sdx[y] = sdy
	}
	return sdx
}

func type18() {
	pic.Show(Pic)
}