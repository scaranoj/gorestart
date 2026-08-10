package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		//close the channel after the loop finishes
		// the send-to channel op has the responsibility to close the channel

		close(c)
	}()

	_, ok := <-c
	fmt.Println("Channel open?", ok)

	// receive
	for v := range c {
		fmt.Println(v)
		//not necessary since you're passing v into println
		// <-c
	}

	//fmt.Println("program about to end")
	fmt.Println()

	_, OK := <-c
	fmt.Println("Channel still open?:", OK)
}
