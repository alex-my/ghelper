package graceful

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/alex-my/ghelper/logger"
	"github.com/alex-my/ghelper/time"
)

// gracefulServer 管理 http.Server
type gracefulServer struct {
	servers []*Server

	// shutdownTimeout 退出时的超时时间，单位: 秒
	shutdownTimeout int

	// restartSignals 响应重启的信号
	restartSignals []os.Signal

	// closeSignals 响应关闭的信号
	closeSignals []os.Signal

	// logEnable 是否开启日志，默认 false
	logEnable bool

	// log 日志
	log logger.Log
}

var (
	gserver                *gracefulServer
	defaultShutdownTimeout = 5
)

func init() {
	gserver = &gracefulServer{
		shutdownTimeout: defaultShutdownTimeout,
	}
}

// isLogEnable 是否启用了日志
func (gs *gracefulServer) isLogEnable() bool {
	return gs.logEnable
}

// logger 获取日志
func (gs *gracefulServer) logger() logger.Log {
	return gs.log
}

func (gs *gracefulServer) updateServerNames() {
	// 如果就一个服务器
	if len(gs.servers) == 1 && gs.servers[0].name == "" {
		gs.servers[0].name = fmt.Sprintf("server-%d", os.Getpid())
		return
	}

	// 给未命名的服务器命名
	for index, server := range gs.servers {
		if server.name == "" {
			server.name = fmt.Sprintf("server-%d", index)
		}
	}
}

func (gs *gracefulServer) close() {
	if len(gs.servers) == 0 {
		return
	}

	logEnable := gs.isLogEnable()
	logger := gs.logger()

	for _, server := range gs.servers {
		if err := server.Close(); logEnable {
			if err != nil && err != http.ErrServerClosed {
				logger.Errorf("Server: %s close, err: %s", server.name, err.Error())
			} else {
				logger.Infof("Server: %s exiting", server.name)
			}
		}
	}
}

func (gs *gracefulServer) shutdown(timeout int) {
	if len(gs.servers) == 0 {
		return
	}

	logEnable := gs.isLogEnable()
	logger := gs.logger()
	for _, server := range gs.servers {
		_applyTimeout := timeout

		// 判断该服务器是否单独指定了超时时间
		if server.timeout != -1 {
			_applyTimeout = server.timeout
		}

		if _applyTimeout > 0 {
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second(_applyTimeout))
			defer cancel()

			if err := server.Shutdown(ctx); logEnable {
				if err != nil && err != http.ErrServerClosed {
					logger.Errorf("Server: %s shutdown with timeout, err: %s", server.name, err.Error())
				} else {
					logger.Infof("Server: %s shutdown with timeout", server.name)
				}
			}
			select {
			case <-ctx.Done():
				if logEnable {
					logger.Infof("server: %s timeout of %d seconds", server.name, _applyTimeout)
				}
			}
		} else {
			ctx := context.TODO()
			if err := server.Shutdown(ctx); logEnable {
				if err != nil && err != http.ErrServerClosed {
					logger.Errorf("Server: %s shutdown, err: %s", server.name, err.Error())
				} else {
					logger.Infof("Server: %s shutdown", server.name)
				}
			}
		}
		if logEnable {
			logger.Infof("Server: %s exiting", server.name)
		}
	}
	if logEnable {
		logger.Infof("All server exiting")
	}
}

func (gs *gracefulServer) restart() {
	// 新实例启动，等父进程(旧实例)退出后，新实例由 init 进程托管

	// 旧实例停止监听，并优雅关闭
	gs.shutdown(gs.shutdownTimeout)
}

func (gs *gracefulServer) listenSignal(f func()) {
	if len(gs.servers) == 0 {
		return
	}

	gs.updateServerNames()

	if len(gs.restartSignals) == 0 {
		gs.restartSignals = []os.Signal{syscall.SIGUSR1, syscall.SIGUSR2}
	}
	if len(gs.closeSignals) == 0 {
		gs.closeSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}

	// 检测信号是否同时出现在 重启信号和关闭信号中
	// 信号数量不多，直接遍历
	for _, sig := range gs.restartSignals {
		for _, sig2 := range gs.closeSignals {
			if sig == sig2 {
				panic(fmt.Sprintf("Sig: %d exist in both the restart and shutdown queues", sig))
			}
		}
	}

	go f()

	sigs := []os.Signal{}

	sigs = append(sigs, gs.restartSignals...)
	sigs = append(sigs, gs.closeSignals...)

	ch := make(chan os.Signal)
	signal.Notify(ch, sigs...)
	sig := <-ch
	signal.Stop(ch)

	logEnable := gs.isLogEnable()
	logger := gs.logger()

	if logEnable {
		logger.Infof("Received signal, sig: %+v", sig)
	}

	// 判断是否是重启服务器
	for _, s := range gs.restartSignals {
		if sig == s {
			if logEnable {
				logger.Info("Restart signal .. restart server ..")
			}
			gs.restart()
			return
		}
	}
	// 判断是否是关闭服务器
	for _, s := range gs.closeSignals {
		if sig == s {
			if logEnable {
				logger.Info("Close signal .. shutdown server ..")
			}
			gs.shutdown(gserver.shutdownTimeout)
			return
		}
	}
}

// AddServer 服务器加进来后，就可以优雅的关闭和重启
func AddServer(server *Server, opts ...Option) {
	fmt.Println("[AddServer] enter")
	if server == nil {
		fmt.Println("[AddServer] server is nil")
	}
	// 添加额外选项
	for _, opt := range opts {
		opt(server)
	}

	gserver.servers = append(gserver.servers, server)
}

// RegisterRestartSignal 设置响应重启的信号，接收到信号之后会优雅的重启服务器
// 如果没有设置，会默认使用 SIGUSR1，SIGUSR2
//
// sig: 重启信号
func RegisterRestartSignal(sig ...os.Signal) {
	if len(sig) == 0 {
		return
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
		return
	}
	if len(gserver.servers) == 0 {
		panic("Please call AddServer first.")
	}

	gserver.closeSignals = append(gserver.closeSignals, sig...)
}

// SetShutdownTimeout 设置优雅退出超时时间
// 服务器会每隔500毫秒检查一次连接是否都断开处理完毕
// 如果超过超时时间，就不再检查，直接退出
// 如果要单独给指定的服务器设置 超时时间，可以使用 WithTimeout
//
// timeout: 单位：秒，当 <= 0 时无效，直接退出
func SetShutdownTimeout(timeout int) {
	gserver.shutdownTimeout = timeout
}

// RegisterShutdownHandler 注册关闭函数
// 按照注册的顺序调用这些函数
// 所有已经添加的服务器都会响应这个函数
// 如果要单独给指定的服务器添加 关闭函数，可以使用 WithShutdownHandler
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
