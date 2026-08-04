// This is successful because the make func creates a chan with a buffer of two
// allowing the value (e.g. 42) to sit on the channel until it's taken off (i.e. read)

package main

import "fmt"

func main() {

	ch := make(chan int, 2)
	ch <- 42
	//No deadlock here since the channel allows 2 values. This runs
	ch <- 43

	//this will pull off 42 while 43 sits in there
	fmt.Println(<-ch)
	//this will pull off 43
	fmt.Println(<-ch)
}
