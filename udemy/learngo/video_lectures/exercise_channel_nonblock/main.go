// this runs. No deadlock as the goroutine holds the concurrent channel write
// until the read from the channel via fmt.Println from func main goroutine runs
// concurrently
package main

import "fmt"

func main() {

	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	//can't just pass `ch` otherwise you'll get the address
	//prefix with the `<-` to indicate you want to "read" the value of ch
	fmt.Println(<-ch)
}
