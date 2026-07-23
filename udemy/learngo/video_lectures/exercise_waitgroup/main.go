package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Arch", runtime.GOARCH)
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	//just add the "go" keyword before your func and you have concurrency, but NOT parallism unless you have multiple cpu/cores
	//The go keyword sends foo off to process while the program keeps running through its sequence
	//foo iss off and running, our control flow doesn't have to wait for foo() to run. So func main() launches one goroutine
	// and foo launches another one
	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	wg.Wait()

}

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}

}
