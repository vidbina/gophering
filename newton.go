package main

import "fmt"
import "math"

func Sqrt(x float64) float64 {
	z := x / 2
	for j := 0; j < 10; j++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	numbers := []float64{64, 4, 1, 100, 25000}
	fmt.Printf("got numbers %v", numbers)
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("\n√%v = %v which should be %v", numbers[i], Sqrt(numbers[i]), math.Sqrt(numbers[i]))
	}
}
