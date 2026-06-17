package concurrency

import (
	"fmt"
	"time"
)

type taskResult struct {
	value int
	err   error
}

func asyncCompute(n int, delayMs int, callback func(int) (int, error)) <-chan taskResult {
	delay := time.Duration(delayMs) * time.Millisecond
	delay = max(delay, 1*time.Millisecond)

	out := make(chan taskResult, 1)
	go func() {
		defer close(out)
		time.Sleep(delay)
		value, err := callback(n)
		if err != nil {
			out <- taskResult{err: err}
			return
		}
		out <- taskResult{value: value}
	}()
	return out
}

func runAsyncTasks() {
	squareCallback := func(n int) (int, error) {
		if n < 0 {
			return 0, fmt.Errorf("negative input: %d", n)
		}
		return n * n, nil
	}

	futures := make([]<-chan taskResult, 4)
	for i, n := range []int{3, 5, -1, 7} {
		futures[i] = asyncCompute(n, 100, squareCallback)
	}

	for i, future := range futures {
		res := <-future
		if res.err != nil {
			fmt.Printf("  task %d error: %v\n", i+1, res.err)
		} else {
			fmt.Printf("  task %d result: %d\n", i+1, res.value)
		}
	}
}
