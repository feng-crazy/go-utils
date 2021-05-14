package workpool

import (
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/sirupsen/logrus"
	"gopkg.in/eapache/queue.v1"
)

// Queue queue
type Queue struct {
	sync.Mutex
	popable *sync.Cond
	buffer  *queue.Queue
	closed  bool
	count   int32
	max     int32
}

// New 创建
func NewQueue(max int32) *Queue {
	ch := &Queue{
		buffer: queue.New(),
	}
	ch.popable = sync.NewCond(&ch.Mutex)
	ch.max = max
	return ch
}

// Pop 取出队列,（阻塞模式）
func (q *Queue) Pop() (v interface{}) {
	c := q.popable
	buffer := q.buffer

	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	for q.Len() == 0 && !q.closed {
		c.Wait()
	}

	if q.closed { // 已关闭
		return
	}

	if q.Len() > 0 {
		v = buffer.Peek()
		buffer.Remove()
		atomic.AddInt32(&q.count, -1)
	}
	return
}

// 试着取出队列（非阻塞模式）返回ok == false 表示空
func (q *Queue) TryPop() (v interface{}, ok bool) {
	buffer := q.buffer

	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	if q.Len() > 0 {
		v = buffer.Peek()
		buffer.Remove()
		atomic.AddInt32(&q.count, -1)
		ok = true
	} else if q.closed {
		ok = true
	}

	return
}

// 插入队列，非阻塞
func (q *Queue) Push(v interface{}) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	if !q.closed {
		if q.Len() == q.max {
			logrus.Error("Queue is full, return")
			return
		}
		q.buffer.Add(v)
		atomic.AddInt32(&q.count, 1)
		q.popable.Signal()
	}
}

// 获取队列长度
func (q *Queue) Len() int32 {
	return atomic.LoadInt32(&q.count)
}

// Close Queue
// After close, Pop will return nil without block, and TryPop will return v=nil, ok=True
func (q *Queue) Close() {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	if !q.closed {
		q.closed = true
		atomic.StoreInt32(&q.count, 0)
		q.popable.Broadcast() // 广播
	}
}

// Wait 等待队列消费完成
func (q *Queue) Wait() {
	for {
		if q.closed || q.Len() == 0 {
			break
		}

		runtime.Gosched() // 出让时间片
	}
}
