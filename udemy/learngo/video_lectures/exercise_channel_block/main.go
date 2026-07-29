// deadlock - this does not run
package main

import "fmt"

func main() {

	ch := make(chan int)

	ch <- 42
	fmt.Println(ch)
}
