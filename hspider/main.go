// Copyright 2022 Baidu Inc. All rights reserved.
// Use of this source code is governed by a xxx
// license that can be found in the LICENSE file.

// Package main is special.  It defines a
// standalone executable program, not a library.
// Within package main the function main is also
// special—it’s where execution of the program begins.
// Whatever main does is what the program does.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"icode.baidu.com/baidu/goodcoder/hdf-mini-spider/config"

	"github.com/baidu/go-lib/log"
	"github.com/baidu/go-lib/log/log4go"

	"icode.baidu.com/baidu/goodcoder/hdf-mini-spider/workpool"
)

const (
	Version            = "v1.0"
	SpiderConfFileName = "spider.conf"
)

var (
	confPath = flag.String("c", "../conf", "root path of configuration")
	logPath  = flag.String("l", "../log", "dir path of log")
	help     = flag.Bool("h", false, "to show help")
	version  = flag.Bool("v", false, "to show version")
)

func Exit(code int) {
	log.Logger.Close()
	time.Sleep(100 * time.Millisecond)
	os.Exit(code)
}

func initLog(logSwitch string, logPath *string, stdOut bool) error {
	log4go.SetLogBufferLength(10000)
	log4go.SetLogWithBlocking(false)

	err := log.Init("mini_spider", logSwitch, *logPath, stdOut, "midnight", 5)
	if err != nil {
		return fmt.Errorf("err in log.Init(): %s", err.Error())
	}

	return nil
}

func main() {
	var err error

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	if *version {
		fmt.Println(Version)
		return
	}

	err = initLog("INFO", logPath, true)
	if err != nil {
		fmt.Printf("initLog(): %s\n", err.Error())
		Exit(-1)
	}

	cfg, err := config.ConfigLoad(filepath.Join(*confPath, SpiderConfFileName))
	if err != nil {
		log.Logger.Error("loader.ConfigLoad(): %s", err.Error())
		Exit(-1)
	}

	rootUrls, err := config.RootUrlRead(cfg.UrlListFile)
	if err != nil {
		log.Logger.Error("loader.RootUrlRead(): %s", err.Error())
		Exit(-1)
	}

	miniSpider := workpool.NewWorkPool()
	miniSpider.Init(cfg, rootUrls)
	miniSpider.Start()
	miniSpider.Wait()

	log.Logger.Info("all tasks done")

	Exit(0)
}
