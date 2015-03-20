package main

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("====")
	fmt.Println(r)
	fmt.Fprint(w, "served")
}

func main() {
	var s Server
	fmt.Println("starting webserver")
	err := http.ListenAndServe("localhost:6666", &s)
	fmt.Println("listening")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("no error")
}
