package workpool

import (
	"regexp"
	"sync"
	"time"

	"github.com/baidu/go-lib/log"
	"icode.baidu.com/baidu/goodcoder/hdf-mini-spider/config"
	"icode.baidu.com/baidu/goodcoder/hdf-mini-spider/parser"
)

type WorkPool struct {
	// 任务队列
	TaskQue *Queue
	// url去重表
	UrlTable sync.Map
	// 任务channel，用于判断是否运行完
	TaskChan chan struct{}
	// 用于等待任务结束
	Wg sync.WaitGroup
	// 最大爬取深度
	MaxDepth int
	// 爬取间隔 单位秒
	CrawlInterval int
	// 爬取任务所使用的go routine数
	ThreadCount int
	// 任务通用配置
	TaskCommonCfg *TaskCommonConfig
	// 站点爬取间隔timer表
	TimerTable sync.Map
}

// NewWorkPool Create a new workpool.
func NewWorkPool() *WorkPool {
	return new(WorkPool)
}

// Init Initialize workpool by config.
// Initialize workpool's task queue by seeds.
func (s *WorkPool) Init(config config.Config, rootUrls []string) {
	targetUrlPattern, _ := regexp.Compile(config.TargetUrl)
	taskCommonCfg := &TaskCommonConfig{
		CrawlTimeout:     config.CrawlTimeout,
		OutputDirectory:  config.OutputDirectory,
		TargetUrlPattern: targetUrlPattern,
	}

	// initialize task queue
	s.TaskQue = NewQueue(-1)
	for _, rootUrl := range rootUrls {
		task := &Task{
			Url:       rootUrl,
			Depth:     0,
			CommonCfg: taskCommonCfg,
		}
		s.TaskQue.Push(task)
	}
	s.Wg = sync.WaitGroup{}

	s.TaskChan = make(chan struct{}, config.ThreadCount)

	s.MaxDepth = config.MaxDepth

	s.CrawlInterval = config.CrawlInterval

	s.ThreadCount = config.ThreadCount

	s.TaskCommonCfg = taskCommonCfg
}

// Start to run tasks.
func (s *WorkPool) Start() {
	log.Logger.Info("start to run tasks")
	s.Wg.Add(s.ThreadCount)
	for i := 0; i < s.ThreadCount; i++ {
		threadId := i
		log.Logger.Info("线程:%d 启动", threadId)
		go func(threadId int) {
			for {
				if s.isDone() {
					log.Logger.Info("线程:%d 结束", threadId)
					break
				}
				task, ok := s.TaskQue.TryPop()
				if ok {
					s.TaskChan <- struct{}{}
					s.RunTask(task.(*Task))
					<-s.TaskChan
				} else {
					time.Sleep(10 * time.Millisecond)
				}

			}
			s.Wg.Done()
		}(threadId)
	}
}

func (s *WorkPool) isDone() bool {
	// s.TaskChan==0说明没有任务在运行了
	if s.TaskQue.Len() == 0 && len(s.TaskChan) == 0 {
		return true
	}
	return false
}

func (s *WorkPool) Wait() {
	s.Wg.Wait()
	close(s.TaskChan)
	s.TaskQue.Close()
}

// RunTask Run single task.
func (s *WorkPool) RunTask(task *Task) {
	if task.Depth >= s.MaxDepth {
		return
	}

	// 避免重复抓取
	// LoadOrStore是Go官方提供的sync.Map的一个方法 第一个参数为key 第二个参数为value
	// 如果task.Url已经存在于urlTable中了则返回的ok的值为true 否则ok的值为false并将task.Url加入到urlTable中
	if _, ok := s.UrlTable.LoadOrStore(task.Url, true); ok {
		// 该url的内容正在抓取或者已经抓取过了 直接返回
		return
	}

	// 控制抓取间隔 防止被封禁
	hostName, err := parser.ParseHostName(task.Url)
	if err != nil {
		log.Logger.Error("%s: parser.ParseHostName(): %s", task.Url, err.Error())
		return
	}
	timer, ok := s.TimerTable.LoadOrStore(hostName, time.NewTimer(time.Duration(s.CrawlInterval)*time.Second))
	if ok {
		select {
		case <-timer.(*time.Timer).C:
		}
		timer.(*time.Timer).Reset(time.Duration(s.CrawlInterval) * time.Second)
	}

	log.Logger.Info("start to crawl %s", task.Url)
	urlList, err := task.Run()
	if err != nil {
		log.Logger.Error("%s", err.Error())
		return
	}

	// generate new tasks
	for _, url := range urlList {
		nextTask := &Task{
			Url:       url,
			Depth:     task.Depth + 1,
			CommonCfg: s.TaskCommonCfg,
		}
		s.TaskQue.Push(nextTask)
	}
	log.Logger.Info("task %s done", task.Url)
}
