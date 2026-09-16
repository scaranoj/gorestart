//partial rewrite of kerninghan and donovan's recover example

package main

import "fmt"

func main() {

	Parse("test")

}

//named func

func Parse(input string) (s *string, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	// ...parser...
	return s, err
}
