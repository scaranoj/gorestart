package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5}
	sum := foo(xi...)
	fmt.Println(sum)

	xb := []int{6, 7, 8, 9}
	xiBar := bar(xb)
	fmt.Println(xiBar)

}

func foo(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	t := 0
	for _, v := range xi {
		t += v
	}
	return t
}

/*
*  create a func with the identifier foo that
   *  takes in a variadic parameter of type int
   *  pass in a value of type []int into your func (unfurl the []int)
   *  returns the sum of all values of type int passed in
*  create a func with the identifier bar that
   *  takes in a parameter of type []int
   *  returns the sum of all values of type int passed in
   *  https://go.dev/play/p/AVSkvMuaz0H
*/
