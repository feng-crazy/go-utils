package eventbus

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/feng-crazy/go-utils/slice"
)

type DataEvent struct {
	Data  interface{}
	Topic string
}

const DataChannelQueueMixSize = 2

// DataChannel 是一个能接收 DataEvent 的 channel
type DataChannel struct {
	Channel        chan DataEvent
	ChannelMaxSize int
}

// DataChannelSlice 是一个包含 DataChannels 数据的切片
type DataChannelSlice []DataChannel

// EventBus 存储有关订阅者感兴趣的特定主题的信息
type EventBus struct {
	// topic 订阅者
	Subscribers map[string]DataChannelSlice
	RWLock      sync.RWMutex
	// 通道订阅者
	Publishers map[DataChannel][]string
}

func NewEventBus() *EventBus {
	return &EventBus{
		Subscribers: map[string]DataChannelSlice{},
		RWLock:      sync.RWMutex{},
		Publishers:  map[DataChannel][]string{},
	}
}

// 该通道不能关闭, 在取消订阅之后, 会自动关闭
func NewDataChannel() DataChannel {
	return DataChannel{
		Channel:        make(chan DataEvent, DataChannelQueueMixSize),
		ChannelMaxSize: DataChannelQueueMixSize,
	}
}

// 该通道不能关闭, 在取消订阅之后, 会自动关闭
func NewDataChannelWithSize(size int) DataChannel {
	if size < DataChannelQueueMixSize {
		size = DataChannelQueueMixSize
	}
	return DataChannel{
		Channel:        make(chan DataEvent, size),
		ChannelMaxSize: size,
	}
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.RWLock.RLock()
	defer eb.RWLock.RUnlock()
	if channels, found := eb.Subscribers[topic]; found {
		// 可以这样做是因为切片引用相同的通道，它们是引用传递的
		// 必须创建一个新切片, 因为是闭包传递
		channels := append(DataChannelSlice{}, channels...)
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, dataChannel := range dataChannelSlices {
				// 如果一个通道阻塞,那么该topic的其他通道都会阻塞
				// 如果通道关闭, 该处会报panic
				if len(dataChannel.Channel) >= dataChannel.ChannelMaxSize {
					// 通道满了,就提取一个丢掉
					logrus.Error("eventbus 通道满了,提取一个丢掉", len(dataChannel.Channel))
					_ = <-dataChannel.Channel
				}
				dataChannel.Channel <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
}

func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.RWLock.Lock()
	defer eb.RWLock.Unlock()
	if prev, found := eb.Publishers[ch]; found {
		// 判断订阅是否存在,存在返回
		if slice.ContainsString(prev, topic) {
			return
		}
		eb.Publishers[ch] = append(prev, topic)
	} else {
		eb.Publishers[ch] = append([]string{}, topic)
	}

	if prev, found := eb.Subscribers[topic]; found {
		eb.Subscribers[topic] = append(prev, ch)
	} else {
		eb.Subscribers[topic] = append([]DataChannel{}, ch)
	}
}

// 取消订阅，如果该通道没有订阅任何topic，将自动关闭,请慎用
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
	if topics, found := eb.Publishers[ch]; found {
		for _, t := range topics {
			if t != topic {
				newTopics = append(newTopics, t)
			} else {
				fmt.Println("--------")
			}
		}

		eb.Publishers[ch] = newTopics
	}

	// 如果该通道没有发布者,这关闭该通道
	if topics, found := eb.Publishers[ch]; found {
		if len(topics) == 0 {
			ch.Close()
		}
	} else {
		ch.Close()
	}
}

func (dc DataChannel) Close() {
	close(dc.Channel)
}
