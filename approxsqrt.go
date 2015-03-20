package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for {
		nz := z - ((math.Pow(z, float64(2)) - float64(x)) / 2 * z)
		if math.Abs(nz-z) < 0.1 {
			return nz
		}
		z = nz
	}
	return 0
}

func main() {
	fmt.Println(Sqrt(2))
}
