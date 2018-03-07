package main

import "fmt"

type Friend interface {
	Hello() string
}

func Hi(f Friend) {
	fmt.Println("hello,", f.Hello())
}

// START OMIT
// using a string for custom type
type good string

// satisfy Friend interface
func (g good) Hello() string {
	return "I am " + string(g)
}

func main() {
	var g good = "Gopher"
	Hi(g)
}

// END OMIT
