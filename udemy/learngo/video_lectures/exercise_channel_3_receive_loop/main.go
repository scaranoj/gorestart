// Get the below code running. Use a for range to receive.

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	c := gen()
// 	receive(c)

// 	fmt.Println("about to exit")
// }

// func gen() <-chan int {
// 	c := make(chan int)

// 	for i := 0; i < 100; i++ {
// 		c <- i
// 	}

// 	return c
// }

// Notice that func gen RETURNS a directional send channel
// receive isn't defined.
// Create a named func called receive that creates a channel and returns
// channels/values
// The receive func will need to also be a loop, try a for-range
// Also convert the standard calls in func main to goroutines

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go gen(c)
	receive(c)

	fmt.Println("about to exit")

}

func gen(c chan int) {

	for i := 0; i < 100; i++ {
		c <- i
	}

	//close(c)

}

func receive(c <-chan int) {

	for v := range c {

		fmt.Println(v)
	}

}
