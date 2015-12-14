package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pickupOrder(series []int, ch chan int) {
	fmt.Println("pickup")
	for i, _ := range series {
		fmt.Printf("(waitng to push %d) ", i)
		ch <- i
		fmt.Printf("(pushed %d) ", i)
	}
}

func monitorProgress(ch chan int) {
	for {
		delay := time.Duration(rand.Int63n(10)*100) * time.Millisecond
		val := <-ch
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("(start %d for %dms) ", val, delay/time.Millisecond)
		time.Sleep(delay)
		fmt.Printf("(end %d)\n", val)
	}
}

func main() {
	ch := make(chan int)
	fmt.Println("hi")

	numbers := []int{1, 2, 4, 4, 2, 3, 4, 2, 3, 4, 1, 2}

	go pickupOrder(numbers, ch)
	go monitorProgress(ch)
	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		// time.Sleep(2 * time.Second)
	}
	fmt.Println("bye")

}
