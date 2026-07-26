//Understanding addressability through interfaces (pointer vs. value semantics)

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area:", s.area())
}

func main() {

	c := circle{8}
	//calling info by passing `c` won't work because the interface expects a pointer
	//compiler will only perform implicit conversion (e.g. `&c`) for method calls on an addressable value
	//the concrete value stored inside an interface is not addressable value (base types
	// info(c)
	info(&c)

}
