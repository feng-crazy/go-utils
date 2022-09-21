package workpool

import (
	"os/exec"
	"testing"

	"icode.baidu.com/baidu/goodcoder/hdf-mini-spider/config"
)

func TestNewWorkPool(t *testing.T) {
	workPool := NewWorkPool()
	if workPool == nil {
		t.Errorf("workPool is nil")
		return
	}
}

func TestWorkPool_Start(t *testing.T) {
	mkCmd := exec.Command("/bin/bash", "-c", "mkdir ../output")
	rmCmd := exec.Command("/bin/bash", "-c", "rm -rf ../output")

	err := mkCmd.Start()
	if err != nil {
		t.Errorf("mkCmd.Start(): %s", err.Error())
		return
	}
	err = mkCmd.Wait()
	if err != nil {
		t.Errorf("mkCmd.Wait(): %s", err.Error())
	}
	defer rmCmd.Start()

	workPool := NewWorkPool()
	cfg := config.Config{
		Spider: config.Spider{
			UrlListFile:     "../data/url.data",
			OutputDirectory: "../output",
			MaxDepth:        1,
			CrawlInterval:   1,
			CrawlTimeout:    1,
			TargetUrl:       ".*.(htm|html)$",
			ThreadCount:     8,
		},
	}
	rootUrls := []string{"http://www.baidu.com", "http://www.sina.com"}

	workPool.Init(cfg, rootUrls)
	workPool.Start()
	workPool.Wait()
}

func TestWorkPool_RunTask(t *testing.T) {
	mkCmd := exec.Command("/bin/bash", "-c", "mkdir ../test_output")
	rmCmd := exec.Command("/bin/bash", "-c", "rm -rf ../test_output")
	err := mkCmd.Start()
	if err != nil {
		t.Errorf("mkCmd.Start(): %s", err.Error())
		return
	}
	err = mkCmd.Wait()
	if err != nil {
		t.Errorf("mkCmd.Wait(): %s", err.Error())
		return
	}
	defer rmCmd.Start()

	workPool := NewWorkPool()
	cfg := config.Config{
		Spider: config.Spider{
			UrlListFile:     "../data/url.data",
			OutputDirectory: "../test_output",
			MaxDepth:        2,
			CrawlInterval:   1,
			CrawlTimeout:    1,
			TargetUrl:       ".*.(htm|html)$",
			ThreadCount:     8,
		},
	}
	rootUrls := []string{"http://www.baidu.com", "http://www.sina.com"}
	workPool.Init(cfg, rootUrls)
	task := &Task{
		Url:       "http://www.test.com",
		Depth:     0,
		CommonCfg: workPool.TaskCommonCfg,
	}
	workPool.RunTask(task)
	task.Url = "http://www.baidu.com"
	workPool.RunTask(task)
}
