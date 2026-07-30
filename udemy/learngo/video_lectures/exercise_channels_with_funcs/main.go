package main

import "fmt"

func main() {

	//create the channels here so they're avail to entire program
	c := make(chan int)
	//rc := make(<-chan int)

	// send
	// remember, goroutines are prefixed before CALLs, NOT defs (unless immed invoked)
	// pass in your channel
	go foo(c)
	// receive
	// omit the 'go' keyword on this one, a goroutine here won't give foo a chance to run/send the value
	// if there were several more, a WaitGroup might be ideal
	bar(c)

	fmt.Println("program about to end")
}

// send
// "chan int" is a type, but we want to send to it, general to directional allowed
func foo(c chan<- int) {
	//send something to the channel declared above
	c <- 42
}

// receive
func bar(c <-chan int) {
	//use a fmt.Println since it's a read type func and you'll see the value
	// in the output to know this func ran before the Println at end of main
	fmt.Println(<-c)
}
