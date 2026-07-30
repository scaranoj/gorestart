package main

import (
	"fmt"
	"runtime"
	"sync"
)

var incrementer = 0

const gs = 100

var mu sync.Mutex
var wg sync.WaitGroup

func main() {
	runtime.NumCPU()
	runtime.NumGoroutine()
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			defer mu.Unlock()
			mu.Lock()
			v := incrementer
			//runtime.Gosched()
			v++
			incrementer = v
			fmt.Println("incrementer", incrementer)
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("end value incrementer", incrementer)
}
