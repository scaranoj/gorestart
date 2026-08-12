//Channel exercise. Get this code working using an anon func literal

/*package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}
*/

package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

// Pseudocode spec/ Pre-mental walkthrough before refactoring (solely from memory after being away from chan for a while):

// Initial example doesn't run because there's nothing reading from the channel
// This means you need a `<-c` somewhere, though wasn't immed sure where to place (scope).
// The whole idea of channels is for the sake of safely sharing variable values through concurrency
// An anon func/func literal is required here, which is a closure that allows you to modify variable values of
// the enclosing function.
// So then with that logic, that means you'd want a func literal/anon func to write to the chan,
// and then have an expression somewhere to read the channel.
// I'm thinking the func literal (closure) will write a value to the channel, and then the enclosing func (main) will
// read it by passing it to println. Only needs to run once, no looping necessary on this one.
// So, to get concurrency, you'll want to make the func literal a goroutine that is off and running
// right after you create the channel. The "closure" is effective because the goroutine and its
// enclosing func (main) are being leveraged to access a shared variable (e.g. `c`), and calling println has
// dual effect of taking the value off the channel while printing it to the console for the user.
