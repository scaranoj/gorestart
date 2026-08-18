package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	//put a value on c
	c <- 1
	//is it open?
	v, ok := <-c
	//read from the channel using a println so you can see the output too
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)

}
