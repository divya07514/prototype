package subscriber

type Subscriber struct {
	Id          string
	DataChannel chan interface{}
}

func (s Subscriber) CloseChannel(one *Subscriber) {
	close(one.DataChannel)
}
