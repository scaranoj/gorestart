// This is successful because the make func creates a chan with a buffer of one
// allowing the value (e.g. 42) to sit on the channel until it's taken off (i.e. read)

package main

import "fmt"

func main() {

	ch := make(chan int, 1)
	ch <- 42
	//nope, only room for one in your channel, sorry...blocked (i.e. "fatal error: all goroutines are asleep - deadlock!")
	ch <- 43

	fmt.Println(<-ch)
}
