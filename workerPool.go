// https://go.dev/play/p/Q-lmGh3aPCm

//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs, results := make(chan int, numJobs), make(chan int, numJobs)
	var wg sync.WaitGroup

	for w := 0; w < 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Println("Processed Result:", res)
	}
}
