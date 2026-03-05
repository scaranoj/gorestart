package main

import "fmt"

func main() {

	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	//use the slice operator to create new slices
	a := xi[:5]
	fmt.Println(a)

	b := xi[5:]
	fmt.Println(b)

	c := xi[2:7]
	fmt.Println(c)

	d := xi[1:6]
	fmt.Println(d)

	// for _, v := range xi {
	// 	fmt.Printf("%v is of type %T\n", v, v)
	// }

}

/*Follow these steps:
* start with this slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
* append to that slice this value
* 52
* print out the slice
* in ONE STATEMENT append to that slice these values
	* 53
	* 54
	* 55
* print out the slice
* append to the slice this slice
	* y := []int{56, 57, 58, 59, 60}
* print out the slice
https://go.dev/play/p/lmPDBWxI8Mo
*/
