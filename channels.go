package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		fmt.Printf("summing %d\r\n", v)
		sum += v
	}
	c <- sum
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x)
	fmt.Println(y)
}
