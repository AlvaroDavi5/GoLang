package concurrency

import (
	"fmt"
)

func Start() {
	// NOTE - Goroutines: multiple goroutines running concurrently, synchronized with a WaitGroup
	fmt.Println("--- Goroutines ---")
	routines := []goRoutine{
		createGoroutine("A", 100, func() { fmt.Println("\t\tA: doing work") }),
		createGoroutine("B", 150, func() { fmt.Println("\t\tB: doing work") }),
		createGoroutine("C", 50, func() { fmt.Println("\t\tC: doing work") }),
	}
	runGoroutines(routines)

	// NOTE - Events: Simple in-process event bus using channels and a subscriber map
	fmt.Println("\n--- Events ---")
	const topicName = "order.created"
	events := []string{"order-001", "order-002"}
	runEvents(topicName, events)

}
