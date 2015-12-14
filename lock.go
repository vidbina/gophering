package main

import "fmt"
import "time"

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("waiting")
		ch <- 12
		fmt.Println("done")
	}()

	go func() {
		fmt.Println("waiting long")
		time.Sleep(10 * time.Second)
		ch <- 21
		fmt.Println("long done")
	}()

	fmt.Printf("got %d\n", <-ch)
	fmt.Println("between")
	fmt.Printf("got %d\n", <-ch)
	time.Sleep(2 * time.Second)
	fmt.Println("end")
}
