package main

import (
	"fmt"
	//	"math/rand"
	"time"
)

func sender(ch chan int, done chan bool, t time.Time) {
	for i := 1; ; i++ {
		//		delay := time.Duration(rand.Int63n(10)) * (10 * time.Millisecond)
		//		if false {
		//			time.Sleep(delay - delay)
		//		}
		fmt.Printf("%d composed %dms later\n", i, time.Since(t))
		ch <- i
		// at this stage the message is definitely sent, we have no means of
		// knowing if the other side has already picked up and processed the
		// message
		fmt.Printf("%d sent at most %dms later\n", i, time.Since(t))

		if i > 10 {
			done <- true
		}
	}
}

func receiver(ch chan int, t time.Time) {
	for {
		fmt.Printf("%d received %dms later\n", <-ch, time.Since(t))
	}
}

func main() {
	done := make(chan bool)
	ch := make(chan int, 5)
	t := time.Now()
	go sender(ch, done, t)
	go receiver(ch, t)
	fmt.Println(<-done)
}
