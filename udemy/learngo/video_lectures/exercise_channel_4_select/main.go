// Starting with this code: https://go.dev/play/p/MvL6uamrJP
// Pull the values off the channel using a select statement
// solution: https://play.golang.org/p/FulKBY5JNj

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	q := make(chan int)
// 	c := gen(q)

// 	receive(c, q)

// 	fmt.Println("about to exit")
// }

// func gen(q <-chan int) <-chan int {
// 	c := make(chan int)

// 	for i := 0; i < 100; i++ {
// 		c <- i
// 	}

// 	return c
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	q := make(chan int)
// 	c := gen(q)

// 	receive(c, q)

// 	fmt.Println("about to exit")
// }

// func gen(q <-chan int) <-chan int {
// 	c := make(chan int)

// 	for i := 0; i < 100; i++ {
// 		c <- i
// 	}
// 	close(c)
// 	return c
// 	}

// func receive(c <-chan int, q <-chan int) {

// 	select {
// 	case <-c:
// 	case <-q:
// 	}

// }

//v2

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	q := make(chan int)
// 	c := gen(q)

// 	receive(c, q)

// 	fmt.Println("about to exit")
// }

// func gen(q <-chan int) <-chan int {
// 	c := make(chan int)

// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			c <- i
// 		}

// 	}()
// 	return c
// }

// func receive(c <-chan int, q chan int) {

// 	for i := 0; i < 100; i++ {
// 		select {
// 		case <-c:
// 			fmt.Println(<-c)
// 		case <-q:
// 			fmt.Println(<-q)
// 		}
// 	}
// }

//v3

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		//we're going to push a value on to q
		q <- 1
		close(c)
	}()

	return c
}

func receive(c, q <-chan int) {

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
