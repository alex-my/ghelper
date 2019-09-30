# 说明

封装日志

# 结构

- 日志接口

  ```go
  const (
    DEBUG = iota
    INFO
    WARN
    ERROR
    FATAL
  )

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
  ```

# 实现

- 简单的实现了`Log`接口，仅输出到控制台
- 打印时间，进程，文件名，函数名，函数，并且使用颜色标记日志级别
- 示例:

  ```text
  [2019-04-18 17:01:47.168][25432][server.go:97-core.(*server).Start][INFO] Framework Version: 0.1.0
  [2019-04-18 17:01:47.168][25432][server.go:99-core.(*server).Start][INFO] PID: 25432
  [2019-04-18 17:01:47.168][25432][router_method.go:113-core.(*methodRoute).output][DEBUG] Static, Method: GET, Path: /
  [2019-04-18 17:01:47.168][25432][server.go:117-core.(*server).Start][DEBUG] Listen on: http://127.0.0.1:8090
  ```
