// package acdc asks if you are ready to rock
package acdc

// Sum uses a for loop to range over a slice and add it's elements
// it adds an unlimited number of values of type int
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s

}
