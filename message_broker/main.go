package main

import (
	"fmt"
	broker "prototype/message_broker/broker"
	"prototype/message_broker/model"
	subscriber "prototype/message_broker/subscriber"
	"strconv"
	"time"
)

func main() {

	subOne := getSubscriber("1")
	subTwo := getSubscriber("2")
	broker := broker.NewBroker()

	broker.Subscribe("test", subOne, subTwo)
	broker.Subscribe("match", subOne, subTwo)

	list := make([]*subscriber.Subscriber, 10)
	list = append(list, subOne)

	startReadingMessages(subOne, subTwo)

	for i := 1; i < 11; i++ {
		message := &model.Message{Data: strconv.Itoa(i)}
		broker.Publish("test", message)
	}

	for i := 100; i < 111; i++ {
		message := &model.Message{Data: strconv.Itoa(i)}
		broker.Publish("match", message)
	}

	time.Sleep(time.Second * 2)
	broker.Unsubscribe("test", subOne)
	broker.Publish("test", &model.Message{Data: strconv.Itoa(12)})
	time.Sleep(time.Second)
}

func getSubscriber(id string) *subscriber.Subscriber {
	return &subscriber.Subscriber{
		Id:          id,
		DataChannel: make(chan interface{}),
	}
}

func startReadingMessages(list ...*subscriber.Subscriber) {
	for _, sub := range list {
		go func(s *subscriber.Subscriber) {
			for {
				select {
				case msg, ok := <-s.DataChannel:
					if !ok {
						str := fmt.Sprintf("Channel for subscriber with ID %s is closed", s.Id)
						fmt.Println(str)
						return
					}
					str := fmt.Sprintf("Message for subscriber with ID %s is %s", s.Id, msg)
					fmt.Println(str)
				}
			}
		}(sub)
	}
}
