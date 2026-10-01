package main

import (
	"errors"
	"fmt"
	"log"
)

//optionally create a new error and assign it to a var to use in your named func below (`sqrt`)
//var myNewError = errors.New("Can't process negative number - package var assignment version")

// create struct type
type sqrtError struct {
	lat  string
	long string
	err  error
}

// create Error()  method (which happens to be the Error() interface's method set - i.e. "it implements the error interface")
//this calles Sprintf and passes the lat, long, and err to it's caller

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	//I can understand throwing away the value with a blank identifier, but how can we even call sqrt here? It's not defined yet.
	//A: semantically named funcs go after func main(), the compiler should be fine with it
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		//Few different options available below:
		//return 0, errors.New("Cannot process negative number")
		//OR
		//return 0, fmt.Errorf("Cannot process negative number %v", f)
		//OR
		//var myError = fmt.Errorf("Can't process negative number (local assignment version) %v", f)
		//return 0, myError
		//OR, uncomment package var assignment above (myNewError) and uncomment the below return
		//return 0, myNewError
		//OR, use the error you just created, which is the point of this exercise:
		//return 0, sqrtError{"2.2", "4.4", errors.New("error: unable to use zero or negative number")}
		//OR, can create a value and use it as a sort of substiution instead
		e := errors.New("error: can't use zero or a negative number")
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}

	}
	return 42, nil
}

// see use of structs with error type in standard library:
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go
