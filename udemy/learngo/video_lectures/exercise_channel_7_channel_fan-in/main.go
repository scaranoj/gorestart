// write a program that
// launches 10 goroutines
// each goroutine adds 10 numbers to a channel
// pull the numbers off the channel and print them

package main

import "fmt"

func main() {

	//make the channel here so the goroutine AND the recieving for loop can access it
	// TIP: write everything in main before abstracting it
	c := make(chan int)

	for i := 0; i < 10; i++ {

		go func() {
			for v := 0; v < 10; v++ {
				c <- v
			}
		}()
		// If using for-range, you'd need a chan close, but can't close the chan above or here because
		// you can't close the channel 10x, nor close here because it'd close before the goroutine
		// completes with all the values on the channel, which is a panic (i.e. read close okay, send close panic)
		// close(c)

	}

	for j := 0; j < 100; j++ {
		fmt.Println(<-c)

	}

	// for v := range c {
	// 	fmt.Println(v)

	// }
	fmt.Println("Program about to exit")

}
