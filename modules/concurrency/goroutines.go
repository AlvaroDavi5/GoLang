package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type goRoutine struct {
	id       string
	delay    time.Duration
	function func()
}

func createGoroutine(identifier string, delayMs int, callback func()) goRoutine {
	return goRoutine{
		id:       identifier,
		delay:    time.Duration(delayMs) * time.Millisecond,
		function: callback,
	}
}

func runGoroutines(routines []goRoutine) {
	var wg sync.WaitGroup

	fmt.Println("Starting goroutines...")

	worker := func(routine goRoutine) {
		defer wg.Done()
		fmt.Printf("\tgoroutine %s: started\n", routine.id)
		time.Sleep(routine.delay * 50)
		if routine.function != nil {
			routine.function()
		}
		fmt.Printf("\tgoroutine %s: done\n", routine.id)
	}

	for i := range routines {
		wg.Add(1)
		go worker(routines[i])
	}

	wg.Wait()

	fmt.Println("All goroutines finished")
}
