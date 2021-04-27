package eventbus

import (
	"fmt"
	"sync"

	"github.com/feng-crazy/go-utils/slice"
)

type DataEvent struct {
	Data  interface{}
	Topic string
}

const DataChannelQueueSize = 10

// DataChannel 是一个能接收 DataEvent 的 channel
type DataChannel chan DataEvent

// DataChannelSlice 是一个包含 DataChannels 数据的切片
type DataChannelSlice []DataChannel

// EventBus 存储有关订阅者感兴趣的特定主题的信息
type EventBus struct {
	Subscribers map[string]DataChannelSlice
	RWLock      sync.RWMutex
	Publisher   map[DataChannel][]string
}

func NewEventBus() *EventBus {
	return &EventBus{
		Subscribers: map[string]DataChannelSlice{},
		RWLock:      sync.RWMutex{},
		Publisher:   map[DataChannel][]string{},
	}
}

// 该通道不能关闭, 在取消订阅之后, 会自动关闭
func NewDataChannel() DataChannel {
	return make(DataChannel, DataChannelQueueSize)
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.RWLock.RLock()
	defer eb.RWLock.RUnlock()
	if channels, found := eb.Subscribers[topic]; found {
		// 可以这样做是因为切片引用相同的通道，它们是引用传递的
		// 必须创建一个新切片, 因为是闭包传递
		channels := append(DataChannelSlice{}, channels...)
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				// 如果一个通道阻塞,那么该topic的其他通道都会阻塞
				// 如果通道关闭, 该处会报panic
				if len(ch) == DataChannelQueueSize {
					// 通道满了,就提取一个
					_ = <-ch
				}
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
}

func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.RWLock.Lock()
	defer eb.RWLock.Unlock()
	if prev, found := eb.Publisher[ch]; found {
		// 判断订阅是否存在,存在返回
		if slice.ContainsString(prev, topic) {
			return
		}
		eb.Publisher[ch] = append(prev, topic)
	} else {
		eb.Publisher[ch] = append([]string{}, topic)
	}

	if prev, found := eb.Subscribers[topic]; found {
		eb.Subscribers[topic] = append(prev, ch)
	} else {
		eb.Subscribers[topic] = append([]DataChannel{}, ch)
	}
}

// 不能使用了要取消订阅,之后没有订阅不能再使用该通道
func (eb *EventBus) UnSubscribe(topic string, ch DataChannel) {
	eb.RWLock.Lock()
	defer eb.RWLock.Unlock()
	newDataChannels := make(DataChannelSlice, 0)
	if channels, found := eb.Subscribers[topic]; found {
		for _, channel := range channels {
			if channel != ch {
				newDataChannels = append(newDataChannels, channel)
			}
		}

		eb.Subscribers[topic] = newDataChannels
	}

	newTopics := make([]string, 0)
	if topics, found := eb.Publisher[ch]; found {
		for _, t := range topics {
			if t != topic {
				newTopics = append(newTopics, t)
			} else {
				fmt.Println("--------")
			}
		}

		eb.Publisher[ch] = newTopics
	}

	// 如果该通道没有发布者,这关闭该通道
	if topics, found := eb.Publisher[ch]; found {
		if len(topics) == 0 {
			ch.Close()
		}
	} else {
		ch.Close()
	}
}

func (dc DataChannel) Close() {
	close(dc)
}
