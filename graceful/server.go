package graceful

import (
	"net/http"
	"os"

	"github.com/alex-my/ghelper/logger"
)

// Package 优雅的关闭和重启 http.Server

// AddServer 服务器加进来后，就可以优雅的关闭和重启
func AddServer(servers ...*http.Server) {
	if len(servers) == 0 {
		panic("servers does not be empty")
	}
	gserver.servers = append(gserver.servers, servers...)
}

// RegisterRestartSignal 设置响应重启的信号，接收到信号之后会优雅的重启服务器
// 如果没有设置，会默认使用 SIGUSR1，SIGUSR2
//
// sig: 重启信号
func RegisterRestartSignal(sig ...os.Signal) {
	if len(sig) == 0 {
		panic("If sig is empty, it is better not to call this function.")
	}
	if len(gserver.servers) == 0 {
		panic("Please call AddServer first.")
	}
	gserver.restartSignals = append(gserver.restartSignals, sig...)
}

// RegisterCloseSignal 设置响应关闭的信号，接收到信号之后会优雅的关闭服务器
// 如果没有设置，会默认使用 SIGINT，SIGTERM
//
// sig: 退出信号
func RegisterCloseSignal(sig ...os.Signal) {
	if len(sig) == 0 {
		panic("If sig is empty, it is better not to call this function.")
	}
	if len(gserver.servers) == 0 {
		panic("Please call AddServer first.")
	}
	gserver.closeSignals = append(gserver.closeSignals, sig...)
}

// SetShutdownTimeout 设置优雅退出超时时间
// 服务器会每隔500毫秒检查一次连接是否都断开处理完毕
// 如果超过超时时间，就不再检查，直接退出
//
// timeout: 单位：秒，当 <= 0 时无效，直接退出
func SetShutdownTimeout(timeout int) {
	gserver.shutdownTimeout = timeout
}

// RegisterShutdownHandler 注册关闭函数
// 按照注册的顺序调用这些函数
func RegisterShutdownHandler(f func()) {
	if len(gserver.servers) == 0 {
		panic("Please call AddServer first.")
	}

	for _, server := range gserver.servers {
		server.RegisterOnShutdown(f)
	}
}

// ListenSignal 监听信号，阻塞
func ListenSignal(f func()) {
	gserver.listenSignal(f)
}

// Close 直接关闭服务器
func Close() {
	gserver.close()
}

// Shutdown 优雅关闭服务器
// 关闭监听
// 执行之前注册的关闭函数(RegisterShutdownHandler)，可以用于清理资源等
// 关闭空闲连接，等待激活的连接变为空闲，再关闭它
//
// timeout: 超时时间，不再等待激活连接变为空闲，直接关闭，系统每隔500毫秒检查一次，单位：秒
func Shutdown(timeout int) {
	gserver.shutdown(timeout)
}

// Restart ..
func Restart() {
	gserver.restart()
}

// SetLogger 设置日志，默认无日志
func SetLogger(logger logger.Log) {
	gserver.logEnable = true
	gserver.log = logger
}
