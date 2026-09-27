//just proving that we can create our own user-define/custom error types by simply implementing the error interface

package main

import (
	"fmt"
)

// don't think fields are even necessary? No, they are not, however it might be nice to be
// able to provide additional information for your custom error. You'll know it will need to be
// a string, and the field name can be anything
type customErr struct {
	info string
}

//type customErr implements the builtin.error interface
//so just create a method and attach it to the customErr type, ya?
// This was my first attempt at writing a method that implements customErr but you were confusing it with interface syntax
// The method set in the builtin.error interface is just `Error() string` (see below)
// func (c customErr) cE() {
//	Error() string
//}

// return apparently needs to be named for the return to work here (compiler barks)
func (ce customErr) Error() (s string) {
	// this was my original return, which compiles, but then this method doesn't do anything useful
	// return s
	// it's goal is to return a custom string. We know that errors.New uses sprintf underneath
	// so why not just use that? Forgot to specify a verb, %v is fine - review https://pkg.go.dev/fmt
	return fmt.Sprintf("Error: %v", ce.info)

}

func main() {

	//var v customErr
	//or
	v := customErr{
		info: "Bad input",
	}
	foo(v)

}

// named func foo that takes an input parm that is of type customErr

func foo(c error) {
	//wasn't part of exercise to print but need to think logically what this program is supposed to do
	//noticed in the hint that the below println was added (the func should DO something of course, right?)
	fmt.Println("foo ran\n", c)
	fmt.Printf("%T/n", c)
}
