package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string
type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "String was "+s)
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi "+s.Who)
}

func main() {
	var port = "6666"
	http.Handle("/string", String("Mona"))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
