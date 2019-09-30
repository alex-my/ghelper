package logger

// 封装 logrus 日志
// logrus: https://github.com/Sirupsen/logrus

import (
	"bufio"
	"os"
	"path"
	"runtime"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type logrusLog struct {
	logger *logrus.Logger
}

// NewLogrusLogger 生成 logrus 日志实例
// logName: 日志文件名称
// logPath: 日志文件存储路径
// caller: 是否打印函数信息
// json: 是否使用 JSON 格式打印
// maxAge: 日志保留时间，例如 180*24*time.Hour，保留 180 天
// rotationTime: 日志切割时间，例如 24*time.Hour 每天切割一次
func NewLogrusLogger(logName, logPath string, caller, json bool, maxAge, rotationTime time.Duration) (Logger, error) {
	basePath := path.Join(logName, logPath)

	writer, err := rotatelogs.New(
		basePath+".%Y%m%d.log",
		rotatelogs.WithLinkName(basePath+".log"),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()

	timeFormat := "2006-01-02 15:04:05.000"

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
	}, &logrus.JSONFormatter{
		TimestampFormat: timeFormat,
	})
	logger.AddHook(lfsHook)

	logger.SetLevel(logrus.DebugLevel)
	if json {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: timeFormat,
		})
	}
	// logrus 固定了 runtime.Caller 级别，造成打印出的是本页面的函数，尚未提供接口进行修改
	// 禁用自带的方法，使用 hook 方式添加自定义方法
	// logger.SetReportCaller(caller)
	// callerHook(logger)
	if caller {
		logger.AddHook(newCallerHook())
	}

	return &logrusLog{
		logger: logger,
	}, nil
}

type callerHook struct{}

func newCallerHook() *callerHook {
	return &callerHook{}
}

func (hook *callerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *callerHook) Fire(entry *logrus.Entry) error {
	pc, file, line, ok := runtime.Caller(9)
	if !ok {
		file = "unknown???"
		line = 0
	}
	_, filename := path.Split(file)
	funcName := runtime.FuncForPC(pc).Name()

	entry.Data["pid"] = os.Getpid()
	entry.Data["file"] = filename
	entry.Data["line"] = line
	entry.Data["func"] = funcName
	return nil
}

// Debug ..
func (l *logrusLog) Debug(v ...interface{}) {
	l.logger.Debugln(v...)
}

// Debugf ..
func (l *logrusLog) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

// Info ..
func (l *logrusLog) Info(v ...interface{}) {
	l.logger.Infoln(v...)
}

// Infof ..
func (l *logrusLog) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

// Warn ..
func (l *logrusLog) Warn(v ...interface{}) {
	l.logger.Warnln(v...)
}

// Warnf ..
func (l *logrusLog) Warnf(format string, v ...interface{}) {
	l.logger.Warnf(format, v...)
}

// Error ..
func (l *logrusLog) Error(v ...interface{}) {
	l.logger.Errorln(v...)
}

// Errorf ..
func (l *logrusLog) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}

// Fatal ..
func (l *logrusLog) Fatal(v ...interface{}) {
	l.logger.Fatalln(v...)
}

// Fatalf ..
func (l *logrusLog) Fatalf(format string, v ...interface{}) {
	l.logger.Fatalf(format, v...)
}

// SetEnable 设置日志是否开启
// able: true 开启; false 关闭
func (l *logrusLog) SetEnable(able bool) {

}

// SetPath 设置日志路径
func (l *logrusLog) SetPath(path string) {
	panic("use NewLogrusLogger(.., logPath, ...)")
}

// SetLevel 设置日志响应级别
func (l *logrusLog) SetLevel(level int) {
	logrusLevel := logrus.DebugLevel
	switch level {
	case INFO:
		logrusLevel = logrus.InfoLevel
	case WARN:
		logrusLevel = logrus.WarnLevel
	case ERROR:
		logrusLevel = logrus.ErrorLevel
	case FATAL:
		logrusLevel = logrus.FatalLevel
	}
	l.logger.SetLevel(logrusLevel)
}

// SetConsoleEnable 是否开启控制台日志
func (l *logrusLog) SetConsoleEnable(able bool) {
	if able {
		l.logger.SetOutput(os.Stdout)
	} else {
		src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			panic(err.Error())
		}
		writer := bufio.NewWriter(src)
		l.logger.SetOutput(writer)
	}
}
