package rabbitmq

import (
	"encoding/json"
)

type MQHandler func(*MQContext)

type MQContext struct {
	Request *Message
	Client  *MQ
	App     *MQApp
}

func (c *MQContext) BindJSON(v interface{}) error {
	return json.Unmarshal(c.Request.Data, v)
}
func (c *MQContext) Push(q *Queue, msg *Message) error {
	return c.Client.Push(q, msg)
}
