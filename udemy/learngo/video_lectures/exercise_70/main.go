package main

import "fmt"

func f() func() {
	return func() {
		fmt.Println(42)
	}
}

func main() {
	v := f()
	v()

}
