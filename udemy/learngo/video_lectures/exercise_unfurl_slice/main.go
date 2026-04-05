package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5}
	x := sum(xi...)
	fmt.Println("The sum is: ", x)
}

//btw - the syntax of a func is
// `func (r receiver) identifier (p paramaters) (return(s)) {...}`

func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
	//now that you returned n, you need to catch that at your func call (value needs to go somewhere after you return it)

}
