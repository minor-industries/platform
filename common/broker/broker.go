package broker

import "sync/atomic"

// https://stackoverflow.com/questions/36417199/how-to-broadcast-message-using-channel

type Message interface {
	Name() string
}

type Broker struct {
	stopCh    chan struct{}
	publishCh chan Message
	subCh     chan chan Message
	unsubCh   chan chan Message
	subCount  int32
}

func NewBroker() *Broker {
	return &Broker{
		stopCh:    make(chan struct{}),
		publishCh: make(chan Message, 1),
		subCh:     make(chan chan Message, 1),
		unsubCh:   make(chan chan Message, 1),
	}
}

func (b *Broker) Start() {
	subs := map[chan Message]struct{}{}
	for {
		select {
		case <-b.stopCh:
			return
		case msgCh := <-b.subCh:
			subs[msgCh] = struct{}{}
			atomic.StoreInt32(&b.subCount, int32(len(subs)))
		case msgCh := <-b.unsubCh:
			delete(subs, msgCh)
			atomic.StoreInt32(&b.subCount, int32(len(subs)))
		case msg := <-b.publishCh:
			for msgCh := range subs {
				// msgCh is buffered, use non-blocking send to protect the broker:
				select {
				case msgCh <- msg:
				default:
				}
			}
		}
	}
}

func (b *Broker) Stop() {
	close(b.stopCh)
}

func (b *Broker) Subscribe() chan Message {
	msgCh := make(chan Message, 5)
	b.subCh <- msgCh
	return msgCh
}

func (b *Broker) Unsubscribe(msgCh chan Message) {
	b.unsubCh <- msgCh
}

func (b *Broker) Publish(msg Message) {
	b.publishCh <- msg
}

func (b *Broker) SubCount() int {
	return int(atomic.LoadInt32(&b.subCount))
}

type Publisher interface {
	Publish(msg Message)
}
