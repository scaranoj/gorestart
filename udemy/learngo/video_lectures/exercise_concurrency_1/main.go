// In addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exists

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("num of CPU:", runtime.NumCPU())
	fmt.Println("num of goroutines", runtime.NumGoroutine())

	wg.Add(2)
	go wg1()
	fmt.Println("num of goroutines", runtime.NumGoroutine())
	go wg2()
	wg.Wait()

	fmt.Println("num of CPU:", runtime.NumCPU())
	fmt.Println("num of goroutines", runtime.NumGoroutine())

}

// func wg1() {
// 	defer wg.Done()
// 	fmt.Println("This is the first func")
// }

// func wg2() {
// 	defer wg.Done()
// 	fmt.Println("This is the second func")
// }

func wg1() {
	for i := 0; i <= 10; i++ {
		defer wg.Wait()
		fmt.Println(i)
	}
	wg.Done()

}

func wg2() {
	defer wg.Wait()
	for j := 0; j <= 10; j++ {
		fmt.Println(j)
	}
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	wg.Done()
}
