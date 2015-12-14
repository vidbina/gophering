package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()

	switch time.Friday {
	case today + 0:
		fmt.Println("It's friday")
	case today - 1:
		fmt.Println("Been there yesterday")
	case today + 1:
		fmt.Println("Tomorrow, my friend")
	case today + 2:
		fmt.Println("Two more days hustling")
	case today - 2:
		fmt.Println("Had a good weekend?")
	default:
		fmt.Println("Stop thinking about it")
	}
}
