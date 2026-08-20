// write a program that
// puts 100 numbers to a channel
// pull the numbers off the channel and print them

package main

import "fmt"

func main() {

	// create a channel to hold the 100 numbers
	c := make(chan int)

	//send 100 numbers to the channel. A classic `for` loop is often used in the textbooks
	//since it's a one-time operation and not going to be used again, just make it a func literal/anon func.
	// Don't forget to make the func (not the `for`) a goroutine, and be sure to close it too since a receive loop is waiting

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		//print v, not c
		fmt.Println(v)
	}
}
