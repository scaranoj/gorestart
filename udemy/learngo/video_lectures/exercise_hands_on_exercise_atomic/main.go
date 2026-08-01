package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var incrementer int64

const gs = 100

var wg sync.WaitGroup

func main() {
	runtime.NumCPU()
	runtime.NumGoroutine()
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("end value incrementer", incrementer)
}
