package rabbitmq

import (
	"github.com/streadway/amqp"
)

// Queue 队列
type Queue struct {
	Name       string     // 必须包含前缀标识使用类型 msg. | rpc. | reply. | notify.
	Key        string     // 和交换机绑定时用的Key
	Durable    bool       // 消息代理重启后，队列依旧存在
	AutoDelete bool       // 当最后一个消费者退订后即被删除
	Exclusive  bool       // 只被一个连接（connection）使用，而且当连接关闭后队列即被删除
	NoWait     bool       // 不需要服务器返回
	ReplyTo    *Queue     // rpc 的消息回应道哪个队列
	Args       amqp.Table // 一些消息代理用他来完成类似与TTL的某些额外功能
	IsDeclare  bool       // 是否已定义

	q            *amqp.Queue
	exchange     *Exchange     // 绑定的交换机
	consumerChan chan *Message // 接收该队列数据的通道
}

func (q *Queue) ReplyQueue() string {
	if q.ReplyTo == nil {
		return ""
	}
	return q.ReplyTo.Name
}

func (q *Queue) GetKey() string {
	if q.Key == "" {
		return q.Name
	}
	return q.Key
}
