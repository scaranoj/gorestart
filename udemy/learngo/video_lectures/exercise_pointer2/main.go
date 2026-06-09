package main

import "fmt"

// print out types and values
func main() {

	i := 42
	p := &i
	fmt.Println(i)
	fmt.Println(p)
	i = 54
	p2 := &i
	fmt.Println(p2) //gives a different address
	*p = 31         //notice the use of the assignment operator
	fmt.Println(*p)

	num := 69
	numPtr := &num
	fmt.Printf("The value of numPtr is: %v\n", numPtr)
	fmt.Printf("The type of numPtr is: %T\n", numPtr)

	num2 := 69
	numPtr2 := &num2 //assigning the address of num2 to numPtr2
	fmt.Printf("num2 has a value of %v and a type of %T\n", num2, num2)
	fmt.Printf("numPtr2 has a value of %v and type of %T\n", numPtr2, numPtr2)
	fmt.Printf("*numPtr2 has a value of %v and a type of %T\n", *numPtr2, *numPtr2)

	*numPtr2 = 22
	fmt.Println(num2)

}

//& = "address of"
//* = dereference (i.e. shows the value at the location in memory)
