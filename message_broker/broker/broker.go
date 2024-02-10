package broker

import (
	"prototype/message_broker/model"
	"prototype/message_broker/subscriber"
	"sync"
)

type Broker struct {
	subscribers map[string][]*subscriber.Subscriber
	mutex       sync.Mutex
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string][]*subscriber.Subscriber),
		mutex:       sync.Mutex{},
	}
}

func (b *Broker) Subscribe(topic string, subscriber ...*subscriber.Subscriber) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.subscribers[topic] = append(b.subscribers[topic], subscriber...)
}

func (b *Broker) Unsubscribe(topic string, subscriber *subscriber.Subscriber) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if subscribers, found := b.subscribers[topic]; found {
		for index, sub := range subscribers {
			if sub == subscriber {
				close(sub.DataChannel)
				b.subscribers[topic] = append(subscribers[:index], subscribers[index+1:]...)
			}
		}
	}
}

func (b *Broker) Publish(topic string, message *model.Message) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if subscribers, found := b.subscribers[topic]; found {
		for _, sub := range subscribers {
			sub.DataChannel <- message.Data
		}
	}
}
