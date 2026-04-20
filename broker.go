package main

import (
	"sync"
)

// Broker represents an in-memory message broker.
type Broker struct {
	messages    []string
	subscribers map[string]*Subscriber
	mu          sync.RWMutex
}

// NewBroker creates a new Broker.
func NewBroker() *Broker {
	return &Broker{
		messages:    make([]string, 0),
		subscribers: make(map[string]*Subscriber),
	}
}

// Publish adds a message to the broker and notifies all active subscribers.
func (b *Broker) Publish(msg string) {
	b.mu.Lock()
	b.messages = append(b.messages, msg)
	
	// Create a copy of subscribers to avoid holding the lock during processing
	subs := make([]*Subscriber, 0, len(b.subscribers))
	for _, sub := range b.subscribers {
		subs = append(subs, sub)
	}
	b.mu.Unlock()

	// Notify all subscribers to process the new message
	for _, sub := range subs {
		sub.Process()
	}
}

// Subscribe registers a subscriber with the broker.
// Requirement: Subscribers MUST subscribe only ONCE; subscribing more than once is a no-op.
func (b *Broker) Subscribe(sub *Subscriber) {
	b.mu.Lock()
	if _, exists := b.subscribers[sub.ID]; exists {
		b.mu.Unlock()
		return
	}
	b.subscribers[sub.ID] = sub
	b.mu.Unlock()

	// Immediately process any existing messages in the broker
	sub.Process()
}
