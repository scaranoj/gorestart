//write a program that performs fan-in from multiple channels to just one

package main

import (
	"fmt"
	"sync"
)

func main() {

	// make your channels
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	//call them as goroutines

	go send(even, odd)

	go receive(even, odd, fanin)

	//print the values as they hit the fanin channel
	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("Program about to exit")

}

//send channel ops

func send(e, o chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

//receive channel
//let's use anon func literals, one to receive values from even/odd and one to fan-in to a single fanin channel
//also will want to use a waitgroup wait here since we're calling it from main as a goroutine and thus
//may finish before send func sends (writes) all of the values

func receive(e, o <-chan int, f chan<- int) {

	//wait groups? I'm thinking yes, esp since you're call to the receive func is in a goroutine

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			//would like to throw away v but compiler won't let me, unless I use an assignment operator? or a classic for loop?
			// just write the value to fanin. Gotta use v and blank identifer not allowed in for-range
			//no sense printing the values now, just print them from main for entire fanin channel
			f <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range o {
			f <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(f)

}
