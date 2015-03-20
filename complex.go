package main

import (
	"fmt"
)

func complexify(x int) {
	const cx = (x) + 2i
	rx, sx := x+1, x
	fmt.Printf("cx is %v, rx is %v and sx is %v", cx, rx, sx)
}

func main() {
	complexify(12)
}
