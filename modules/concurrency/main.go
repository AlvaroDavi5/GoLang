package concurrency

import (
	"fmt"
)

func Start() {
	// ANCHOR - Goroutines: multiple goroutines running concurrently, synchronized with a WaitGroup.
	fmt.Println("--- Goroutines ---")
	routines := []goRoutine{
		createGoroutine("A", 100, func() { fmt.Println("\t\tA: doing work") }),
		createGoroutine("B", 150, func() { fmt.Println("\t\tB: doing work") }),
		createGoroutine("C", 50, func() { fmt.Println("\t\tC: doing work") }),
	}
	runGoroutines(routines)
}
