package main

import (
	"fmt"
	"sync"
)

// START OMIT
var wg sync.WaitGroup

func inc(x int, c chan int) {
	fmt.Println("incrementing:", x)
	c <- x + 1
	wg.Done()
}

func main() {
	cc := make(chan int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go inc(i, cc)
	}
	fmt.Println("collect results")
	go func() {
		for x := range cc {
			fmt.Println("got back:", x)
		}
	}()
	wg.Wait()
	close(cc)
}

// END OMIT
