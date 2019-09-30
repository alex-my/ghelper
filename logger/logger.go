package logger

// Package 日志接口，自定义日志需要实现本接口

// 日志级别
const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

// Log 日志接口
type Log interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	// Fatal 最终调用 panic
	Fatal(v ...interface{})
	// Fatalf 最终调用 panic
	Fatalf(format string, v ...interface{})

	// Enable 设置日志是否开启
	// able: true 开启; false 关闭
	Enable(able bool)

	// SetPath 设置日志路径
	SetPath(path string)

	// SetLevel 设置日志响应级别
	SetLevel(level int)

	// SetConsoleEnable 是否开启控制台日志
	SetConsoleEnable(able bool)
}