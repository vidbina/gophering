package main

import (
	"code.google.com/p/go-tour/pic"
	"math"
)

func draw(a int, b int) uint8 {
	//return uint8(a*b/2)
	return uint8(a ^ b)
	//return uint8(math.Pow(float64(a), float64(b)))
	//return uint8(math.Pow(float64(b), float64(a)))
}
func Pic(dx, dy int) [][]uint8 {
	res := [][]uint8{}
	for i := range make([]uint8, dy) {
		col := []uint8{}
		for j := range make([]uint8, dx) {
			col = append(col, draw(i, j)) //uint8(i*j/2))
		}
		res = append(res, col)
	}
	return res
}

func main() {
	pic.Show(Pic)
}
