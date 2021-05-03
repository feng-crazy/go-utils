package rpcpool

import (
	"errors"
	"net/rpc"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var ErrShutdown = rpc.ErrShutdown

const constTPC = "tcp"

type RPCPool interface {
	GetAddr() string
	Call(serviceMethod string, args interface{}, reply interface{}) error
	Close() error
}

type rpcPool struct {
	isClosed        bool
	addr            string
	pool            innerPool
	lock            *sync.Mutex
	conns           int
	minConns        int
	maxConns        int
	idleConnTimeout time.Duration
	waitConnTimeout time.Duration
	clearPeriod     time.Duration
}

func NewRPCPool(addr string, minConns, maxConns int, idleConnTimeout, waitConnTimeout,
	clearPeriod time.Duration) (RPCPool, error) {
	p := &rpcPool{
		addr:            addr,
		pool:            make(innerPool, maxConns),
		lock:            &sync.Mutex{},
		minConns:        minConns,
		maxConns:        maxConns,
		idleConnTimeout: idleConnTimeout,
		waitConnTimeout: waitConnTimeout,
		clearPeriod:     clearPeriod,
	}
	err := p.setupPool()
	if err != nil {
		return nil, err
	}
	go p.startPoolManager()
	return p, nil
}

func (p *rpcPool) GetAddr() string {
	return p.addr
}

func (p *rpcPool) Call(serviceMethod string, args interface{}, reply interface{}) error {
	c := p.poolGet()
	defer p.poolPut(c)
	c.activeTime = time.Now()
	return c.call(serviceMethod, args, reply)
}

func (p *rpcPool) Close() error {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.isClosed = true
	for p.conns > 0 {
		c := p.poolGet()
		err := c.close()
		if err != nil {
			return err
		}
		p.conns--
	}
	return nil
}

func (p *rpcPool) newRPCClient() (*rpcClient, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.conns == p.maxConns {
		return nil, errors.New("rpc pool full")
	}
	c, err := p.pool.new(p.addr)
	if err != nil {
		return nil, err
	}
	p.conns++
	return c, nil
}

func (p *rpcPool) setupPool() error {
	for i := 0; i < p.minConns; i++ {
		c, err := p.newRPCClient()
		if err != nil {
			return err
		}
		p.poolPut(c)
	}
	return nil
}

func (p *rpcPool) poolPut(c *rpcClient) {
	p.pool <- c
}

func (p *rpcPool) poolGet() *rpcClient {
	for {
		select {
		case c := <-p.pool:
			return c
		case <-time.After(p.waitConnTimeout):
			c, err := p.newRPCClient()
			if err == nil {
				return c
			}
		}
	}
}

func (p *rpcPool) startPoolManager() {
	ticker := time.NewTicker(p.clearPeriod)
	for range ticker.C {
		if p.isClosed {
			break
		}
		p.clearPool()
	}
}

func (p *rpcPool) clearPool() {
	p.lock.Lock()
	defer p.lock.Unlock()
	conns := p.conns
	bucket := make(innerPool, p.maxConns)
	for i := 0; i < p.conns && p.minConns < conns; i++ {
		var c *rpcClient
		check := true
		select {
		case c = <-p.pool:
			if time.Since(c.activeTime) > p.idleConnTimeout {
				conns--
				err := c.close()
				if err != nil {
					log.Error(err)
				}
				c = nil
			}
		case <-time.After(p.waitConnTimeout):
			check = false
		}
		if !check {
			break
		}
		if c != nil {
			bucket <- c
		}
	}
	for len(bucket) > 0 {
		p.poolPut(<-bucket)
	}
	p.conns = conns
}
