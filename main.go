package main

import (
	"fmt"
)

func main() {
	// 1. Create Broker
	broker := NewBroker()

	// 2. Create Publisher
	publisher := NewPublisher(broker)

	// 3. for 1 in 5: Publish() // publish distinct messages
	fmt.Println("--- Publishing first 5 messages ---")
	for i := 1; i <= 5; i++ {
		publisher.Publish(fmt.Sprintf("Message %d", i))
	}

	// 4. Create Subscriber1
	sub1 := NewSubscriber("Subscriber1", broker)

	// 5. Subscriber1.Subscribe()
	fmt.Println("\n--- Subscriber1 Subscribing ---")
	sub1.Subscribe()

	// 6. for 1 in 5: Publish() // publish distinct messages
	fmt.Println("\n--- Publishing next 5 messages ---")
	for i := 6; i <= 10; i++ {
		publisher.Publish(fmt.Sprintf("Message %d", i))
	}

	// 7. Create Subscriber2
	sub2 := NewSubscriber("Subscriber2", broker)

	// 8. Subscriber2.Subscribe()
	fmt.Println("\n--- Subscriber2 Subscribing ---")
	sub2.Subscribe()

	// Final verification note:
	// 10 messages should be printed out from both subscriber 1 and 2 totalling 20 message logs
}
