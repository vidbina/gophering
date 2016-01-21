package main

import "fmt"

func successful_channel() {
	c := make(chan int)
	go func(ch chan int) {
		fmt.Println("Sent 42")
		ch <- 42
	}(c)
	fmt.Printf("Received %v\r\n", <-c)
}

func send_and_receive_in_goroutine() {
	c := make(chan int)
	go func(ch chan int) {
		ch <- 42
		fmt.Println("Sent 42")
	}(c)
	go func(ch chan int) {
		fmt.Printf("Received %v\r\n", <-ch)
	}(c)
	fmt.Println("Never rets as goroutines are nonblocking and the program exits")
	fmt.Println("We need the receive to block until the message arrives")
}

func nothing_to_receive() {
	// problematic because we trying to receive something but nothing is sent
	c := make(chan int)
	fmt.Printf("Received %v\r\n", <-c)
}

func sending_into_emptiness() {
	c := make(chan int)
	c <- 42
	fmt.Println("Sent 42")
}

func receive_blocking_send() {
	c := make(chan int)
	fmt.Printf("Received %v\r\n", <-c)
	c <- 42
}

func send_failed_because_of_nonreceiving_channel() {
	c := make(chan int)
	c <- 42
	fmt.Printf("Received %v\r\n", <-c)
}

func main() {
	successful_channel()
	send_and_receive_in_goroutine()
	//sending_into_emptiness()
	//nothing_to_receive()
	//receive_blocking_send() // deadlock
	//send_failed_because_of_nonreceiving_channel() // deadlock
}
