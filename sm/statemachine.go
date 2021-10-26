package sm

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type StateMachineState = string                   // 状态
type StateMachineHandler func() StateMachineState // 处理方法，并返回新的状态

type StateMachine struct {
	mu       sync.Mutex                                // 排他锁
	state    StateMachineState                         // 当前状态
	handlers map[StateMachineState]StateMachineHandler // 每一个状态都可以出发有限个事件，执行有限个处理

	tick   time.Duration
	ticker *time.Ticker

	stopC  chan struct{}
	exited bool
}

// 获取当前状态
func (sm *StateMachine) GetState() StateMachineState {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.state
}

// 设置当前状态
func (sm *StateMachine) SetState(newState StateMachineState) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.state = newState
}

// 某状态添加事件处理方法
func (sm *StateMachine) AddHandler(state StateMachineState, handler StateMachineHandler) *StateMachine {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if sm.handlers == nil {
		sm.handlers = make(map[StateMachineState]StateMachineHandler)
	}

	if _, ok := sm.handlers[state]; !ok {
		sm.handlers[state] = handler
	}

	return sm
}

// 事件处理
func (sm *StateMachine) Call() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	handle := sm.handlers[sm.state]
	if handle == nil {
		logrus.Error("状态处理函数没有添加")
		return
	}
	state := handle()
	sm.state = state
	return
}

func (sm *StateMachine) StartWithTicker() {
	sm.exited = false
	sm.ticker = time.NewTicker(sm.tick)

	go func() {
		for !sm.exited {
			select {
			case <-sm.ticker.C:
				sm.Call()
			case <-sm.stopC:
				sm.exited = true
				sm.ticker.Stop()
				return
			}
		}
	}()
}

func (sm *StateMachine) StopWithTicker() {
	sm.stopC <- struct{}{}
}

func (sm *StateMachine) Start() {
	go func() {
		for {
			sm.mu.Lock()
			if sm.exited {
				sm.mu.Unlock()
				return
			}
			sm.mu.Unlock()
			sm.Call()
			time.Sleep(sm.tick)
		}

	}()
}

func (sm *StateMachine) Stop() {
	sm.mu.Lock()
	sm.exited = true
	sm.mu.Unlock()
}

// 实例化
func NewStateMachine(initState StateMachineState, tick time.Duration) *StateMachine {
	return &StateMachine{
		state:    initState,
		handlers: make(map[StateMachineState]StateMachineHandler),
		tick:     tick,
		stopC:    make(chan struct{}),
	}
}
