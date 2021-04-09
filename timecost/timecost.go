package timecost

import (
	"fmt"
	"time"
)

var TimeCostMap map[string]time.Time

func init() {
	TimeCostMap = make(map[string]time.Time)
}

func Since(name string) {
	TimeCostMap[name] = time.Now()
}

func Until(name string) {
	if v, ok := TimeCostMap[name]; ok {
		tc := time.Since(v)
		fmt.Printf("%s time cost = %v\n", name, tc)
	} else {
		fmt.Printf("请先使用since\n")
	}

}

func timeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("time cost = %v\n", tc)
	}
}
