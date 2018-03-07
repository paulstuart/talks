package main

import "fmt"

// START OMIT
var c = make(chan int)

func moar(n int) {
	n++
	c <- n
}

func main() {
	go moar(23)
	n := <-c
	fmt.Println("the answer is:", n)
}

// END OMIT
