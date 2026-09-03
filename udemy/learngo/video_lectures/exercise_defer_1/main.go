package main

import "fmt"

func main() {
	// notice the value of i is 0 and not 1. This is because i is evaluated as 0 when
	// the deferred call is registered, while the call itself executes later, as `a` returns.
	// Evaluation happens when encountered by the runtime, the call happens later when `a` returns.

	a()

}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
}
