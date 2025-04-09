// https://go.dev/play/p/5Cst-DiLSBg

//go:build ignore

package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(time.Second * 3)
	ch <- "Response from Server 1"
}

func server2(ch chan string) {
	time.Sleep(time.Second * 6)
	ch <- "Response from Server 2"
}

func main() {
	res1, res2 := make(chan string), make(chan string)
	go server1(res1)
	go server2(res2)

	select {
	case msg := <-res1:
		fmt.Println(msg)
	case msg := <-res2:
		fmt.Println(msg)
	}
}
