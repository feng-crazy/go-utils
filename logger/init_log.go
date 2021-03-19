package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var logOnce sync.Once

func InitLogs(logPath, projectName, level string) {
	logOnce.Do(func() {
		// 日志级别
		InitLogLevel(level)
		// 使用自定义格式
		logFormatter := new(LogFormatter)
		logFormatter.Init()
		logrus.SetFormatter(logFormatter)
		// 开启行号
		logrus.SetReportCaller(true)
		/*
			    日志轮转相关函数
			    `WithLinkName` 为最新的日志建立软连接
			    `WithRotationTime` 设置日志分割的时间，隔多久分割一次
			    `WithMaxAge` 和 `WithRotationCount` 二者只能设置一个
				`WithMaxAge` 设置文件清理前的最长保存时间
			    `WithRotationCount` 设置文件清理前最多保存的个数
				`WithRotationSize` 设置文件大小切分单位byte
		*/
		logFilePath := filepath.Join(logPath, projectName+"%Y%m%d.log")
		logFileLinkPath := filepath.Join(logPath, projectName+".log")
		writer, _ := rotatelogs.New(
			logFilePath,
			rotatelogs.WithLinkName(logFileLinkPath),
			rotatelogs.WithMaxAge(15*24*time.Hour),
			rotatelogs.WithRotationTime(24*time.Hour),
			rotatelogs.WithRotationSize(200*1000*1000),
		)
		// logrus.SetOutput(writer)
		logrus.SetOutput(io.MultiWriter(os.Stdout, writer)) // 控制台和文件打印
	})
}

func InitLogLevel(level string) {
	switch level {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	default:
		logrus.Fatal("log conf only allow [info, debug, warn], please check your confguire")
	}
}

// 日志自定义格式
type LogFormatter struct {
	ginRe *regexp.Regexp
}

func (s *LogFormatter) Init() {
	s.ginRe = regexp.MustCompile("(?m)[\r\n]+^.*gin-gonic.*$")
}

func LogStack() string {
	pc := make([]uintptr, 10)
	n := runtime.Callers(9, pc)
	frames := runtime.CallersFrames(pc[:n])

	var frame runtime.Frame
	more := n > 0
	output := ""
	for more {
		frame, more = frames.Next()
		output = output + fmt.Sprintf("%s:%d \n %s\n", frame.File, frame.Line, frame.Function)
	}
	return output
}

func (s *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("[2006-01-02|15:04:05.000]")
	var file string
	var l int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		l = entry.Caller.Line
	}

	msg := fmt.Sprintf("%s[%s:%d][GID:%d][%s]: %s\n", timestamp, file, l, getGID(), strings.ToUpper(entry.Level.String()), entry.Message)
	if entry.Level <= logrus.ErrorLevel {
		stackInfo := LogStack()

		stackInfo = s.ginRe.ReplaceAllString(stackInfo, "")

		msg = msg + stackInfo
	}

	return []byte(msg), nil
}

// 获取当前协程ID
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
