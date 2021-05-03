package rpcpool

import (
	"net/rpc"
	"time"
)

type rpcClient struct {
	client     *rpc.Client
	activeTime time.Time
}

func (c *rpcClient) call(serviceMethod string, args interface{}, reply interface{}) error {
	return c.client.Call(serviceMethod, args, reply)
}

func (c *rpcClient) close() error {
	return c.client.Close()
}
