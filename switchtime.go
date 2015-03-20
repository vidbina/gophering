package main

import (
	"fmt"
	"time"
)

func main() {
	switch {
	case time.Hour < 12:
		fmt.Print("good morning")
		fallthrough
	case time.Hour < 17:
		fmt.Print("good afternoon")
	case time.Hour > 0 && time.Hour < 4:
		fmt.Print("what an unholy time")
	default:
		fmt.Print("poop")
	}
}
