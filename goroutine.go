// https://go.dev/play/p/2ffUoAdT3O9

//go:build ignore

package main

import (
	"fmt"
	"sync"
)

func printNumber(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("The number is:", n)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go printNumber(i, &wg)
	}

	wg.Wait()
}
