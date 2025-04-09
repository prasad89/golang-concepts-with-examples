// https://go.dev/play/p/yI8ii88fT2T

//go:build ignore

package main

import "fmt"

func fib(n int, ch chan<- int) {
	defer close(ch)

	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
}

func main() {
	n := 10
	ch := make(chan int, n)

	go fib(n, ch)

	for num := range ch {
		fmt.Println(num)
	}
}
