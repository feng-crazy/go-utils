package eventbus

import "sync"

type DataEvent struct {
	Data  interface{}
	Topic string
}

// DataChannel 是一个能接收 DataEvent 的 channel
type DataChannel chan DataEvent

// DataChannelSlice 是一个包含 DataChannels 数据的切片
type DataChannelSlice []DataChannel

// EventBus 存储有关订阅者感兴趣的特定主题的信息
type EventBus struct {
	Subscribers map[string]DataChannelSlice
	RWLock      sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		Subscribers: map[string]DataChannelSlice{},
		RWLock:      sync.RWMutex{},
	}
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.RWLock.RLock()
	if channels, found := eb.Subscribers[topic]; found {
		// 这样做是因为切片引用相同的数组，即使它们是按值传递的
		// 因此我们正在使用我们的元素创建一个新切片，从而正确地保持锁定
		channels := append(DataChannelSlice{}, channels...)
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.RWLock.RUnlock()
}

func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.RWLock.Lock()
	if prev, found := eb.Subscribers[topic]; found {
		eb.Subscribers[topic] = append(prev, ch)
	} else {
		eb.Subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.RWLock.Unlock()
}
