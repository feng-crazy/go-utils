package sm

import (
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	InitDriverState     = StateMachineState("初始化")
	RegisterDriverState = StateMachineState("注册")
	SyncChannelState    = StateMachineState("同步通道")
	SyncDeviceState     = StateMachineState("同步设备")
	WaitConnectState    = StateMachineState("等待连接状态")
	RunState            = StateMachineState("正常运行")
	StopState           = StateMachineState("停止运行")

	InitDriverHandle = func() StateMachineState {
		logrus.Info("InitDriverHandle...")
		return WaitConnectState
	}

	WaitConnectHandle = func() StateMachineState {
		logrus.Info("WaitConnectHandle...")
		time.Sleep(1 * time.Second)
		return RegisterDriverState
	}

	RegisterDriverHandle = func() StateMachineState {
		logrus.Info("RegisterDriverHandle...")
		return SyncChannelState
	}

	SyncChannelHandle = func() StateMachineState {
		logrus.Info("SyncChannelHandle...")
		return SyncDeviceState
	}

	SyncDeviceHandle = func() StateMachineState {
		logrus.Info("SyncDeviceHandle...")
		return RunState
	}

	RunHandle = func() StateMachineState {
		logrus.Info("RunHandle...")
		return RunState
	}

	StopHandle = func() StateMachineState {
		logrus.Info("StopHandle...")
		return StopState
	}
)

func TestStateMachine(t *testing.T) {
	sm := NewStateMachine(InitDriverState, 1*time.Second)

	sm.AddHandler(InitDriverState, InitDriverHandle)
	sm.AddHandler(RegisterDriverState, RegisterDriverHandle)
	sm.AddHandler(SyncChannelState, SyncChannelHandle)
	sm.AddHandler(SyncDeviceState, SyncDeviceHandle)
	sm.AddHandler(WaitConnectState, WaitConnectHandle)
	sm.AddHandler(RunState, RunHandle)
	sm.AddHandler(StopState, StopHandle)

	sm.Start()
	select {
	case <-time.After(10 * time.Second):
		sm.SetState(StopState)
	}

	time.Sleep(3 * time.Second)
	sm.Stop()

	time.Sleep(1 * time.Second)
	fmt.Println("exit")
}
