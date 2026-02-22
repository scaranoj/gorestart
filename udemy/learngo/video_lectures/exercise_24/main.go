package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("x is %v\n", x)

	if x < 100 {
		fmt.Printf("x is %v, which is betw 0 and 100\n", x)
	} else if x >= 101 && x <= 200 {
		fmt.Printf("x is %v, which is betw 101 and 200\n", x)
	} else if x >= 201 && x <= 250 {
		fmt.Printf("x is %v, which is betw 201 and 250\n", x)
	}

	//inc0 := rand.Intn(0)
	inc1 := rand.Intn(1)
	inc2 := rand.Intn(2)
	inc3 := rand.Intn(3)
	//fmt.Printf("0 input to Intn func is %v\n", inc0)
	fmt.Printf("Inputting number 1 to the Intn func is %v\n", inc1)
	fmt.Printf("Inputting number 2 to the Intn func is %v\n", inc2)
	fmt.Printf("Inputting number 3 to the Intn func is %v\n", inc3)
}

/*re: inclusive, exclusive – does rand.Intn()
○ include zero in its output? a: yes
○ include 250 in its output? a: no
○ show this in code using the numbers 0 - 3
*/
