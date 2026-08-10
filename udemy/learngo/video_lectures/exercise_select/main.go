// show how a select statement can pull values off of a channel
// it can also send but not in this exercise
// just send values to a channel depending on whether their odd or even
// use a channel for signaling the completion of the send and read funcs (call the channel quit and place at end of send loop)

package main

import "fmt"

func main() {

	//I'm guessing we want to make a few channels first

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//We'll call two seperate functions. Let's call one named send, and one named receive, which will accept directional channels
	//The send call will be from main to the send func while running concurrently in a new goroutine.
	//The receive call is a direct call from main
	//Receive is runnning from main so it'll block main until receive returns

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("about to exit")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0

}

func receive(e, o, q <-chan int) {
	//infinite loop
	for {
		select {
		//select says: "Hey, which of these channels can I pull a value off of?"
		case v := <-e:
			fmt.Println("read from the even channel:", v)
		case v := <-o:
			fmt.Println("read from the odd channel:", v)
		case v := <-q:
			fmt.Println("read from the quit channel:", v)
			return
		}
	}
}
