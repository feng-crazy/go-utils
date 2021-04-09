package timecost

import (
	"testing"
	"time"
)

func TestTimeCost(t *testing.T) {
	// 注意，是对 timeCost()返回的函数进行调用，因此需要加两对小括号
	defer timeCost()()
	total := 0
	for i := 1; i <= 10; i++ {
		total += i
		time.Sleep(1 * time.Millisecond)
	}
}

func TestTimeCost1(t *testing.T) {
	// 注意，是对 timeCost()返回的函数进行调用，因此需要加两对小括号
	Since("test")
	total := 0
	for i := 1; i <= 10; i++ {
		total += i
		time.Sleep(1 * time.Millisecond)
	}
	Until("test")
}
