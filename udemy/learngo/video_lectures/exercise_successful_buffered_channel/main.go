// This is successful because the make func creates a chan with a buffer of one
// allowing the value (e.g. 42) to sit on the channel until it's taken off (i.e. read)

package main

import "fmt"

func main() {

	ch := make(chan int, 1)
	ch <- 42
	//can't just pass `ch` otherwise you'll get the address since channels are reference types
	//prefix with the `<-` to indicate you want to "read" the value of ch
	fmt.Println(<-ch)
}
