package main

import (
	"fmt"
	"math"
)

func main() {

	s1 := Square{length: 5, width: 5}
	//fmt.Println(s1.Area(5, 5))
	fmt.Println(Info(s1))
	c1 := Circle{radius: 5}
	fmt.Println(Info(c1))
	//fmt.Println(c1.Area(5))
}

type Square struct {
	length float64
	width  float64
}
type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func (s Square) Area() float64 {
	a := s.length * s.width
	return a
}
func (c Circle) Area() float64 {
	a := math.Pi * math.Pow(c.radius, 2)
	return a
}

// create a func INFO which takes type shape and then prints the area
func Info(s Shape) float64 {
	return s.Area()
}

/*
  * create a type SQUARE
    * length float64
    * width float64
  * create a type CIRCLE
    * radius float64
  * attach a method to each that calculates AREA and returns the area
    * circle area= π r 2
      * math.Pi
      * math.Pow
    * square area = L * W
  * create a type SHAPE that defines an interface as anything that has the AREA method
  * create a func INFO which takes type shape and then prints the area
  * create a value of type square
  * create a value of type circle
  * use func info to print the area of square
  * use func info to print the area of circle
  * https://go.dev/play/p/8BFxl2GXgcw
*/
