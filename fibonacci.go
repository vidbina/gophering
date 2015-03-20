package main

import "fmt"

func fib(n int) int {
	switch n {
	case 1, 2:
		return 1
	case 0:
		return 0
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("fib(%d) = %d\n", i, fib(i))
	}
}
