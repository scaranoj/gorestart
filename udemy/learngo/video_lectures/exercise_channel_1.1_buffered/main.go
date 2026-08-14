//Channel exercise. Get this code working simply by using a buffered channel

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

func main() {

	c := make(chan int, 1)

	c <- 42

	//fmt.Println(<-c)
}

// The original commented code above deadlocks right away since it's
// an unbuffered channel, but will run when it's buffered
// since you're effectively telling the channel to hold on to the value (i.e. baton)
// until the receiver (reader) is ready to pull it off the channel
