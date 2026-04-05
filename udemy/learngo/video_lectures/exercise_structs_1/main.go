package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
		city string
	}

	//an empty struct literal
	bob := person{}

	//the first struct literal style
	lydia := person{
		"lydia",
		32,
		"Denver",
	}

	//the second struct literal style
	jason := person{
		name: "jason",
		age:  32,
		city: "Denver",
	}

	fmt.Println(jason)
	fmt.Println(lydia)
	fmt.Println(bob)
	fmt.Printf("%T \t %#v\n", jason, jason)
	fmt.Println(jason.age, jason.city, jason.name)
}
