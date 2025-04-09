// https://go.dev/play/p/AnTcfrDKYNQ

//go:build ignore

package main

import (
	"fmt"
	"net/http"
)

func isLiveOrNot(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

func main() {
	url := "http://www.google.com"
	fmt.Println(isLiveOrNot(url))
}
