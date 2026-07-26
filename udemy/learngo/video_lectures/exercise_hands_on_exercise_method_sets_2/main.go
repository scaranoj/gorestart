package main

import "fmt"

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("Hola!")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	b := person{"bob"}
	// won't work because the interface expects a pointer value
	// also compiler can't do implicit address conversion because interface values aren't addressable
	// saySomething(b)
	// this works as the value passed satisfies the interface, which has a pointer receiver method set
	saySomething(&b)

}
