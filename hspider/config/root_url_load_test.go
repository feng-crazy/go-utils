package config

import (
	"testing"
)

func TestRootUrlRead(t *testing.T) {
	// 测试正常情况
	result, err := RootUrlRead("../data/url.data")
	if err != nil {
		t.Errorf("../data/url.data is valid but there is an error")
		return
	}
	t.Log(result)
}
