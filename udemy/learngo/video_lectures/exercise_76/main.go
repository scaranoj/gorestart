//  Correct the code to do the following:
//   * implement the "youngin" interface
//     * you can change anything in the code to do so except the interface
//   * maintain consistency within the code
//     * receivers should stick with either POINTER or VALUE semantics
//   * when choosing between POINTER or VALUE semantics, understand why you chose one or the other

package main

import "fmt"

// uncomment this code to run it

type dog struct {
	first string
}

func (d *dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

func (d *dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()
	// youngRun(d1) // doesn't run

	d2 := &dog{"Padget"}
	d2.walk()
	d2.run()
	youngRun(d2)

}
