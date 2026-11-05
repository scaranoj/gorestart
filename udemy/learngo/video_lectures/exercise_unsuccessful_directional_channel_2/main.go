package main

import "fmt"

func main() {

	//the `<-` syntax indicates the channel is directional
	//since the arrow syntax is right of the channel, it's a channel that only allows sending values to it
	c := make(chan<- int, 2)
	c <- 42
	c <- 43
	//these reads won't work because c is a send-only type channel, it only allows incoming values sent to it - send-only
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-------")
	fmt.Printf("%T\n", c)

}
