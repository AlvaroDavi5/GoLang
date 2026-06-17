package concurrency

import (
	"fmt"
	"sync"
)

func runChannels() {
	// Unbuffered channel – sender blocks until receiver reads.
	unbuffered := make(chan string)
	go func() { unbuffered <- "hello from unbuffered channel" }()
	fmt.Println(" ", <-unbuffered)

	// Buffered channel – sender does not block while buffer has space.
	buffered := make(chan int, 3)
	buffered <- 10
	buffered <- 20
	buffered <- 30
	close(buffered)
	for v := range buffered {
		fmt.Printf("  buffered value: %d\n", v)
	}

	// Fan-out / fan-in: distribute work across goroutines and collect results.
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	fanOut := func(in <-chan int, out chan<- int) {
		for n := range in {
			out <- n * n
		}
	}

	var wg sync.WaitGroup
	for w := 0; w < 3; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fanOut(jobs, results)
		}()
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("  fan-in result: %d\n", r)
	}
}
