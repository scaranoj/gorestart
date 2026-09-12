package main

import "fmt"

func main() {

	div60()

}

//named func

func div60() {

	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}

		//don't need the index
		for _, val := range []int{10, 50, 0, 100} {
			fmt.Println(val)
		}
	}()

}
