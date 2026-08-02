package main

import "fmt"

func main() {

	//again, the `<-` syntax indicates the channel is directional
	//since the arrow syntax is left of the channel, it's a channel that only allows reading/taking values from it - receive-only
	c := make(<-chan int, 2)
	//these writes won't work because c is a recieve-only type channel, it only allows values to be read/taken from it - receive-only
	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-------")
	fmt.Printf("%T\n", c)

}
