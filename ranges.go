package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i, v := range pow {
		fmt.Println(i, v)
	}
}
