package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("runtime:", runtime.GOOS)
	fmt.Println("arch:", runtime.GOARCH)

}
