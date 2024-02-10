package subscriber

// Subscriber One subscriber can read from only one data source
type Subscriber struct {
	Id          string
	DataChannel chan interface{}
}

func (s Subscriber) CloseChannel(one *Subscriber) {
	close(one.DataChannel)
}
