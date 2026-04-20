package main

import (
	"fmt"
	"sync"
)

// Subscriber reads messages from the broker.
type Subscriber struct {
	ID           string
	broker       *Broker
	readPosition int
	mu           sync.Mutex
}

// NewSubscriber creates a new Subscriber.
func NewSubscriber(id string, b *Broker) *Subscriber {
	return &Subscriber{
		ID:     id,
		broker: b,
	}
}

// Subscribe registers the subscriber to the broker.
func (s *Subscriber) Subscribe() {
	s.broker.Subscribe(s)
}

// Process reads messages from the broker starting from the last read position.
func (s *Subscriber) Process() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.broker.mu.RLock()
	defer s.broker.mu.RUnlock()

	for s.readPosition < len(s.broker.messages) {
		msg := s.broker.messages[s.readPosition]
		fmt.Printf("[%s] Received: %s\n", s.ID, msg)
		s.readPosition++
	}
}
