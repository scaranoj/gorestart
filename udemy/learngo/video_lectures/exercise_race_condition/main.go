// The shared variable causing the race condition is counter.
// All the goroutines read and write counter concurrently:
//     v := counter (read)
//     counter = v (write)
// There’s no synchronization (e.g., a mutex or atomic operations) around those
// accesses, the goroutines can interleave and overwrite each other, producing a data
// race and an incorrect final counter value.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("count", counter)

}
