package main

import "fmt"

func main() {
	s := make([]byte, 5)
	fmt.Printf("A %T with len %v and cap %v\n", s, len(s), cap(s))

	t := s[2:4]
	fmt.Printf("Sliced from 2 through 4 has len %v and cap %v\n", len(t), cap(t))

	u := s[1:3]
	fmt.Printf("Sliced from 1 through 3 has len %v and cap %v\n", len(u), cap(u))

	v := s[0:3]
	fmt.Printf("Sliced from 0 through 3 has len %v and cap %v\n", len(v), cap(v))

	w := s[3:4]
	fmt.Printf("Slived from 3 through 4 has len %v and cap %v\n", len(w), cap(w))

	x := s[3:]
	fmt.Printf("Slived from 3 through N has len %v and cap %v\n", len(x), cap(x))

	y := s[:2]
	fmt.Printf("Slived from 0 through 2 has len %v and cap %v\n", len(y), cap(y))
}
