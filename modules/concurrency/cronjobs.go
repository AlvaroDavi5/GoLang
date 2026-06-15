package concurrency

import (
	"fmt"
	"time"
)

type cronJob struct {
	id       string
	interval time.Duration
	maxTicks int
	job      func()
}

func createCronjob(identifier string, intervalMs int, maxTicks int, function func()) cronJob {
	return cronJob{
		id:       identifier,
		maxTicks: maxTicks,
		interval: time.Duration(intervalMs) * time.Millisecond,
		job:      function,
	}
}

func (cron *cronJob) start() {
	ticker := time.NewTicker(cron.interval)
	defer ticker.Stop()

	done := make(chan struct{})
	count := 0

	go func() {
		for {
			select {
			case t := <-ticker.C:
				count++
				fmt.Printf("\tCronjob '%s' tick %d at %s\n", cron.id, count, t.Format("15:04:05.000"))
				cron.job()
				if count >= cron.maxTicks {
					close(done)
					return
				}
			case <-done:
				return
			}
		}
	}()

	<-done
	fmt.Printf("Cronjob '%s' finished after %d ticks\n", cron.id, cron.maxTicks)
}
