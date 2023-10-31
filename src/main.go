package main

import "fmt"

func main() {
	s := newServer(withId("asdf"), withMaxConn(25))
	fmt.Printf("%+v\n", s)
}
