//Get this code running

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	cs := make(chan<- int)

// 	go func() {
// 		cs <- 42
// 	}()
// 	fmt.Println(<-cs)

// 	fmt.Printf("------\n")
// 	fmt.Printf("cs\t%T\n", cs)
// }

//Answer below:

//Simplest would be to just change the new chan declaration to general by removing the directional syntax (`<-`)

package main

import (
	"fmt"
)

func main() {
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}

//other option would be to create funcs that accept directional chans as arguments
//the code below will run successfully

// package main

//directional channels are usually seen as func parameter types in func declarations
//one design pattern is to just create a non-dir channel and then pass the channel as an argument to func paramaters
//as a pipeline of handoffs

// import (
// 	"fmt"
// )

// func foo(c chan<- int) {
// 	c <- 42
// }
// func bar(c <-chan int) {
// 	fmt.Println(<-c)
// }

// func main() {
// 	cs := make(chan int)

// 	go foo(cs)
// 	bar(cs)

// 	fmt.Printf("------\n")
// 	fmt.Printf("cs\t%T\n", cs)
// }
