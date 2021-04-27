package eventbus

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func printDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}

func publisTo(eb *EventBus, topic string, data string) {
	for {
		eb.Publish(topic, data)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func TestEventBus(t *testing.T) {
	var eb = NewEventBus()

	ch1 := NewDataChannel()
	ch2 := NewDataChannel()
	ch3 := NewDataChannel()

	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)
	eb.Subscribe("topic2", ch3)

	go publisTo(eb, "topic1", "Hi topic 1")
	go publisTo(eb, "topic2", "Welcome to topic 2")

	eb.UnSubscribe("topic2", ch3)
	// eb.UnSubscribe("topic2", ch2)
	// eb.UnSubscribe("topic1", ch1)
	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch2", d)
			// case d := <-ch3:
			// 	go printDataEvent("ch3", d)
		}
	}
}

func TestEventBus1(t *testing.T) {
	var eb = NewEventBus()

	ch1 := make(chan DataEvent)

	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch1)
	eb.Subscribe("topic3", ch1)

	go publisTo(eb, "topic1", "Hi topic 1")
	go publisTo(eb, "topic2", "Welcome to topic 2")
	go publisTo(eb, "topic3", "This is topic 3")

	// eb.UnSubscribe("topic3", ch1)
	for event := range ch1 {
		fmt.Println(event.Topic)
		fmt.Println(event.Data)
		fmt.Println("-----------------")
	}
}

func TestEventBus2(t *testing.T) {
	var eb = NewEventBus()
	ch1 := make(chan DataEvent, 0)

	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch1)
	eb.Subscribe("topic3", ch1)

	eb.Publish("topic1", "Hi topic 1")
	eb.Publish("topic2", "Welcome to topic 2")
	eb.Publish("topic3", "This is topic 3")

	eb.Publish("topic1", "Hi topic 1")
	eb.Publish("topic2", "Welcome to topic 2")
	eb.Publish("topic3", "This is topic 3")

	go func() {
		for event := range ch1 {
			fmt.Println(event.Topic)
			fmt.Println(event.Data)
			fmt.Println("-----------------")
		}
	}()

	time.Sleep(2 * time.Second)
}

func TestEventBus3(t *testing.T) {
	var eb = NewEventBus()

	ch1 := NewDataChannel()
	ch2 := NewDataChannel()
	ch3 := NewDataChannel()

	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)
	eb.Subscribe("topic3", ch3)

	go publisTo(eb, "topic1", "Hi topic 1")
	go publisTo(eb, "topic2", "Welcome to topic 2")

	eb.UnSubscribe("topic3", ch3)
	// eb.UnSubscribe("topic2", ch2)
	// eb.UnSubscribe("topic1", ch1)

	for {
		time.Sleep(1 * time.Second)
	}
}

func TestEventBus4(t *testing.T) {
	var eb = NewEventBus()

	ch1 := NewDataChannel()

	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch1)
	eb.Subscribe("topic2", ch1)

	go publisTo(eb, "topic1", "Hi topic 1")
	go publisTo(eb, "topic2", "Welcome to topic 2")

	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		}
	}

	// for {
	// 	time.Sleep(1 * time.Second)
	// }
}
