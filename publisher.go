package main

// Publisher is responsible for writing messages to the broker.
type Publisher struct {
	broker *Broker
}

// NewPublisher creates a new Publisher.
func NewPublisher(b *Broker) *Publisher {
	return &Publisher{broker: b}
}

// Publish sends a message to the broker.
func (p *Publisher) Publish(msg string) {
	p.broker.Publish(msg)
}
