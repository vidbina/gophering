package main

import "fmt"

type something interface {
	Describe() string
}

type Something struct {
	Name string
}

func (s *Something) Describe() string {
	return fmt.Sprintf("Name is %v", s.Name)
}

/*
func (s *Something) Rename(name string) string {
	s.Name = name
	return s.Name
}
*/

//func return_something() Something
func main() {
	//s := get_something()
	//fmt.Printf("%T is %v\n", s, s)
	//fmt.Println(s.Describe())

	p := get_something_ptr()
	fmt.Println(p.Describe())
	fmt.Printf("%T is %v\n", p, p)
	fmt.Println(p.Describe())
}

// A function returning an interface may simply return an instance of the
// struct or a pointer to it, pointers are preferred for the very same reasons
// why function pointers are used
/*func get_something() something {
	return Something{"Ball"}
}*/

func get_something_ptr() something {
	return &Something{"Star"}
}
