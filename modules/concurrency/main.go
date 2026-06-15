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
	eventsTopic1 := []string{"order-001", "order-002"}
	eventsTopic2 := []string{"order-003", "order-004"}
	subscribersTopic1 := []string{"Alice", "Bob"}
	subscribersTopic2 := []string{"Charlie", "Diana"}
	runEvents("topic1", subscribersTopic1, eventsTopic1)
	runEvents("topic2", subscribersTopic2, eventsTopic2)

	// NOTE - Cronjobs: Ticker-based periodic job that runs a fixed number of times then stops
	fmt.Println("\n--- Cronjobs ---")
	cron1 := createCronjob("Job 1", 2000, 1, func() { fmt.Println("\t\tJob 1: executing task...") })
	cron2 := createCronjob("Job 2", 2000, 2, func() { fmt.Println("\t\tJob 2: executing task...") })
	cron3 := createCronjob("Job 3", 3000, 2, func() { fmt.Println("\t\tJob 3: executing task...") })
	cron4 := createCronjob("Job 4", 2000, 3, func() { fmt.Println("\t\tJob 4: executing task...") })
	cron1.start()
	cron2.start()
	cron3.start()
	cron4.start()
}
