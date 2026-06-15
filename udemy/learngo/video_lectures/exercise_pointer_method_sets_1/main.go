package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Printf("I'm taking", d.first, "for a walk")
}
func (d *dog) run() {
	d.first = "bob"
	fmt.Printf("I'm taking", d.first, "for a run")
}

func main() {

	d1 := dog{"Fido"}
	d1.walk() //remember, when calling a method, don't use the type name, use the variable instance name of the type (e.g. `d1`)
	d1.run()

	sp := &dog{"Sugarpuss"}
	sp.run()
	sp.walk()
}
