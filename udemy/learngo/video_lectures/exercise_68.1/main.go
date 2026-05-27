package main

import "fmt"

// func main() {

// 	func(a, b int) {
// 		s := a + b
// 		fmt.Println(s)
// 	}(2, 2)
// }

func main() {
	// this works but non-idiomatic. You're returing an error because you're using a another func (fmt.Println) to print the output of s (sum)
	// and fmt.Println returns an error. This would be fine but there's no opportunity to check for an error prior to your return. If an error is encounterd
	// from using the println func, it's too late (i.e. end of the function)
	func(a, b int) (int, error) {
		s := a + b
		return fmt.Println(s)
	}(2, 2)
}
