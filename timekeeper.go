package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 6:
		fmt.Println("Hope you are sleeping")
	case t.Hour() < 8:
		fmt.Println("Gymtime")
	case t.Hour() < 12:
		fmt.Println("Hustle")
	case t.Hour() < 18:
		fmt.Println("The hustle continues")
	case t.Hour() < 20:
		fmt.Println("Get social")
	default:
		fmt.Println("Get some rest")
	}
}
