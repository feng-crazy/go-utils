package rpcpool

import (
	"net/rpc"
	"time"
)

type innerPool chan *rpcClient

func (ip innerPool) new(addr string) (*rpcClient, error) {
	client, err := rpc.Dial(constTPC, addr)
	if err != nil {
		return nil, err
	}
	return &rpcClient{
		client:     client,
		activeTime: time.Now(),
	}, nil
}
