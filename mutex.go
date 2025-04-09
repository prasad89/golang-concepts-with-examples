// https://go.dev/play/p/HusoUqO9X69

//go:build ignore

package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
	wg      sync.WaitGroup
)

func incrementCounter() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go incrementCounter()
	}
	wg.Wait()
	fmt.Println("Final Counter Value:", counter)
}
