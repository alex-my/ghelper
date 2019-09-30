package logger

// 简易日志，只输出到终端
// 带日志定位，时间，颜色

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/alex-my/ghelper/time"
)

// eg:
// 输出: 时间、进程号、文件名称、函数名称、日志级别
// [2019-04-18 17:01:47.168][25432][server.go:97-core.(*server).Start][INFO] Framework Version: 0.1.0
// [2019-04-18 17:01:47.168][25432][server.go:99-core.(*server).Start][INFO] PID: 25432
// [2019-04-18 17:01:47.168][25432][router_method.go:113-core.(*methodRoute).output][DEBUG] Static, Method: GET, Path: /
// [2019-04-18 17:01:47.168][25432][server.go:117-core.(*server).Start][DEBUG] Listen on: http://127.0.0.1:8090

const (
	// ColorBlackC 前景 黑色
	ColorBlackC = 30
	// ColorBlackB 背景 黑色
	ColorBlackB = 40

	// ColorRedC 前景 红色
	ColorRedC = 31
	// ColorRedB 背景 红色
	ColorRedB = 41

	// ColorGreenC 前景 绿色
	ColorGreenC = 32
	// ColorGreenB 背景 绿色
	ColorGreenB = 42

	// ColorYellowC 前景 黄色
	ColorYellowC = 33
	// ColorYellowB 背景 黄色
	ColorYellowB = 43

	// ColorBlueC 前景 蓝色
	ColorBlueC = 34
	// ColorBlueB 背景 蓝色
	ColorBlueB = 44

	// ColorBurgundyC 前景 紫红色
	ColorBurgundyC = 35
	// ColorBurgundyB 背景 紫红色
	ColorBurgundyB = 45

	// ColorCyanC 前景 青色
	ColorCyanC = 36
	// ColorCyanB 背景 青色
	ColorCyanB = 46

	// ColorWhiteC 前景 白色
	ColorWhiteC = 37
	// ColorWhiteB 背景 白色
	ColorWhiteB = 47
)

// simpleLog 简单日志，输出到终端
type simpleLog struct {
	// able true: 开启日志，false 关闭日志
	able bool

	// console true: 开启终端日志，false 关闭终端日志
	console bool

	// level 日志级别
	level int
}

// NewLogger ..
func NewLogger() Logger {
	return &simpleLog{able: true, console: true}
}

// Debug ..
func (l *simpleLog) Debug(v ...interface{}) {
	l.log(DEBUG, v...)
}

// Debugf ..
func (l *simpleLog) Debugf(format string, v ...interface{}) {
	l.logf(DEBUG, format, v...)
}

// Info ..
func (l *simpleLog) Info(v ...interface{}) {
	l.log(INFO, v...)
}

// Infof ..
func (l *simpleLog) Infof(format string, v ...interface{}) {
	l.logf(INFO, format, v...)
}

// Warn ..
func (l *simpleLog) Warn(v ...interface{}) {
	l.log(WARN, v...)
}

// Warnf ..
func (l *simpleLog) Warnf(format string, v ...interface{}) {
	l.logf(WARN, format, v...)
}

// Error ..
func (l *simpleLog) Error(v ...interface{}) {
	l.log(ERROR, v...)
}

// Errorf ..
func (l *simpleLog) Errorf(format string, v ...interface{}) {
	l.logf(ERROR, format, v...)
}

// Fatal ..
func (l *simpleLog) Fatal(v ...interface{}) {
	l.log(FATAL, v...)
	panic("")
}

// Fatalf ..
func (l *simpleLog) Fatalf(format string, v ...interface{}) {
	l.logf(FATAL, format, v...)
	panic("")
}

// SetEnable 设置日志是否开启
// able: true 开启; false 关闭
func (l *simpleLog) SetEnable(able bool) {
	l.able = able
}

// SetPath ..
func (l *simpleLog) SetPath(path string) {

}

// SetLevel ..
func (l *simpleLog) SetLevel(level int) {
	l.level = level
}

// SetConsoleEnable ..
func (l *simpleLog) SetConsoleEnable(able bool) {
	// 本日志仅输出到终端
	l.able = able
	l.console = able
}

func levelName(level int) string {
	var name string
	switch level {
	case DEBUG:
		name = "DEBUG"
	case INFO:
		name = "INFO"
	case WARN:
		name = "WARN"
	case ERROR:
		name = "ERROR"
	case FATAL:
		name = "FATAL"
	default:
		name = fmt.Sprintf("UNKNOWN LEVEL: %d", level)
	}

	return name
}

func levelColor(level int) int {
	var color int
	switch level {
	case DEBUG:
		color = ColorCyanC
	case INFO:
		color = ColorGreenC
	case WARN:
		color = ColorYellowC
	case ERROR:
		color = ColorRedC
	case FATAL:
		color = ColorBurgundyC
	default:
		color = ColorBlackC
	}

	return color
}

var bufferPool *sync.Pool

// Buffer 从池中获取 buffer
func Buffer() *bytes.Buffer {
	buff := bufferPool.Get().(*bytes.Buffer)
	buff.Reset()
	return buff
}

// ReleaseBuffer 将 buff 放入池中
func ReleaseBuffer(buff *bytes.Buffer) {
	bufferPool.Put(buff)
}

func runtimeMessage(level int) string {
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "unknown???"
		line = -1
	}
	_, filename := path.Split(file)

	funcName := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	pid := os.Getpid()

	buff := Buffer()
	defer ReleaseBuffer(buff)

	buff.WriteString("[")
	buff.WriteString(time.Date(time.YMDHMSM))
	buff.WriteString("][")
	buff.WriteString(strconv.FormatInt(int64(pid), 10))
	buff.WriteString("][")
	buff.WriteString(filename)
	buff.WriteString(":")
	buff.WriteString(strconv.FormatInt(int64(line), 10))
	buff.WriteString("-")
	buff.WriteString(funcName[len(funcName)-1])
	buff.WriteString("]")

	buff.WriteString("\x1b[")
	buff.WriteString(strconv.Itoa(levelColor(level)))
	buff.WriteString(";1m[")

	buff.WriteString(levelName(level))
	buff.WriteString("]")

	buff.WriteString("\x1b[39;22m")

	return buff.String()
}

func (l *simpleLog) log(level int, v ...interface{}) {
	if l.able && l.console && level >= l.level {
		fmt.Println(runtimeMessage(level), fmt.Sprint(v...))
	}
}

func (l *simpleLog) logf(level int, format string, v ...interface{}) {
	if l.able && l.console && level >= l.level {
		fmt.Println(runtimeMessage(level), fmt.Sprintf(format, v...))
	}
}

func init() {
	bufferPool = &sync.Pool{}
	bufferPool.New = func() interface{} {
		return &bytes.Buffer{}
	}
}
