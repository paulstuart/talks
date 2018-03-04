package main

import "fmt"

// START OMIT
type Friend interface {
	Hello() string
}

func Hi(f Friend) {
	fmt.Println("hello,", f.Hello())
}

// END OMIT

// IN OMIT

// create a custom type based on an integer
type nice int

// satisfy the Friend interface
func (n nice) Hello() string {
	return fmt.Sprintf("my values is %d", n)
}

func main() {
	var n nice = 23
	Hi(n)
}

// OUT OMIT
