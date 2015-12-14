package main

import (
	"fmt"
	"time"
)

func calculate(set []int, i int, ch chan int) {
	sum := 0
	for _, v := range set {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d adds %d to %d, ", i, v, sum)
		sum += v
	}
	ch <- sum // intermediate update
}

func main() {
	numbers := []int{10, 12, 8, 5, 5, 0, 3, 7, 20, 5}

	ch := make(chan int)
	go calculate(numbers[:len(numbers)/2], 1, ch)
	go calculate(numbers[len(numbers)/2:], 2, ch)
	x, y := <-ch, <-ch

	fmt.Println("\n", x, y)
}
