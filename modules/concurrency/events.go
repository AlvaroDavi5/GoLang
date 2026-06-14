package concurrency

import (
	"fmt"
	"sync"
)

type EventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]chan string
}

func (eb *EventBus) subscribeOnTopic(topic string) <-chan string {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	ch := make(chan string, 4)
	eb.subscribers[topic] = append(eb.subscribers[topic], ch)
	return ch
}

func (eb *EventBus) publishOnTopic(topic string, payload string) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	for _, ch := range eb.subscribers[topic] {
		ch <- payload
	}
}

func (eb *EventBus) closeTopic(topic string) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	for _, ch := range eb.subscribers[topic] {
		close(ch)
	}
	delete(eb.subscribers, topic)
}

func newEventBus() *EventBus {
	return &EventBus{subscribers: make(map[string][]chan string)}
}

func runEvents(topic string, subscribers []string, events []string) {
	bus := newEventBus()
	var waitGroup sync.WaitGroup

	for i, subscriber := range subscribers {
		sub := bus.subscribeOnTopic(topic)
		waitGroup.Add(1)

		go func(id int, subscriber string, ch <-chan string) {
			defer waitGroup.Done()
			for event := range ch {
				fmt.Printf("Subscriber %s received: %q\n", subscriber, event)
			}
		}(i, subscriber, sub)
	}

	for _, event := range events {
		bus.publishOnTopic(topic, event)
	}
	bus.closeTopic(topic)

	waitGroup.Wait()
}
