package graceful

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/alex-my/ghelper/logger"
	"github.com/alex-my/ghelper/time"
)

// gracefulServer 管理 http.Server
type gracefulServer struct {
	server *Server

	// shutdownTimeout 退出时的超时时间，单位: 秒
	shutdownTimeout int

	// restartSignals 响应重启的信号
	restartSignals []os.Signal

	// closeSignals 响应关闭的信号
	closeSignals []os.Signal

	// log 日志
	logger logger.Log
}

var (
	gserver                *gracefulServer
	defaultShutdownTimeout = 5
)

func init() {
	gserver = &gracefulServer{shutdownTimeout: defaultShutdownTimeout}
}

// RegisterRestartSignal 设置响应重启的信号，接收到信号之后会优雅的重启服务器
// 如果没有设置，会默认使用 SIGUSR1，SIGUSR2
//
// sig: 重启信号
func RegisterRestartSignal(sig ...os.Signal) error {
	if len(sig) == 0 {
		return errors.New("sig does not be empty")
	}
	if gserver.server == nil {
		return errors.New("use graceful.NewServer first")
	}

	gserver.restartSignals = append(gserver.restartSignals, sig...)
	return nil
}

// RegisterCloseSignal 设置响应关闭的信号，接收到信号之后会优雅的关闭服务器
// 如果没有设置，会默认使用 SIGINT，SIGTERM
//
// sig: 退出信号
func RegisterCloseSignal(sig ...os.Signal) error {
	if len(sig) == 0 {
		return errors.New("sig does not be empty")
	}
	if gserver.server == nil {
		return errors.New("use graceful.NewServer first")
	}

	gserver.closeSignals = append(gserver.closeSignals, sig...)
	return nil
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
func RegisterShutdownHandler(f func()) error {
	if gserver.server == nil {
		return errors.New("use graceful.NewServer first")
	}

	gserver.server.RegisterOnShutdown(f)

	return nil
}

// ListenSignal 监听信号，阻塞
func ListenSignal(f func()) {
	if len(gserver.restartSignals) == 0 {
		gserver.restartSignals = []os.Signal{syscall.SIGUSR1, syscall.SIGUSR2}
	}
	if len(gserver.closeSignals) == 0 {
		gserver.closeSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}

	// 检测信号是否同时出现在 重启信号和关闭信号中
	// 信号数量不多，直接遍历
	for _, sig := range gserver.restartSignals {
		for _, sig2 := range gserver.closeSignals {
			if sig == sig2 {
				gserver.logger.Fatalf("sig: %d exist in both the restart and shutdown queues", sig)
			}
		}
	}

	go f()

	sigs := []os.Signal{}

	sigs = append(sigs, gserver.restartSignals...)
	sigs = append(sigs, gserver.closeSignals...)

	ch := make(chan os.Signal)
	signal.Notify(ch, sigs...)
	sig := <-ch
	signal.Stop(ch)

	gserver.logger.Infof("received signal, sig: %+v", sig)

	// 判断是否是重启服务器
	for _, s := range gserver.restartSignals {
		if sig == s {
			gserver.logger.Info("restart signal .. restart server ..")
			Restart()
			return
		}
	}
	// 判断是否是关闭服务器
	for _, s := range gserver.closeSignals {
		if sig == s {
			gserver.logger.Info("close signal .. shutdown server ..")
			Shutdown()
			return
		}
	}
}

// Close 直接关闭服务器
func Close() {
	err := gserver.server.Close()
	if err != nil && err != http.ErrServerClosed {
		gserver.logger.Errorf("server close, err: %s", err.Error())
	} else {
		gserver.logger.Info("server exiting")
	}
}

// Shutdown 优雅关闭服务器
// 关闭监听
// 执行之前注册的关闭函数(RegisterShutdownHandler)，可以用于清理资源等
// 关闭空闲连接，等待激活的连接变为空闲，再关闭它
//
// timeout: 超时时间，不再等待激活连接变为空闲，直接关闭，系统每隔500毫秒检查一次，单位：秒
func Shutdown() {
	logger := gserver.logger
	server := gserver.server
	timeout := gserver.shutdownTimeout

	if timeout > 0 {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second(timeout))
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil && err != http.ErrServerClosed {
			logger.Errorf("server shutdown with timeout, err: %s", err.Error())
		} else {
			logger.Info("server shutdown with timeout")
		}

		select {
		case <-ctx.Done():
			logger.Infof("server timeout of %d seconds", timeout)
		}
	} else {
		ctx := context.TODO()
		err := server.Shutdown(ctx)

		if err != nil && err != http.ErrServerClosed {
			logger.Errorf("server shutdown, err: %s", err.Error())
		} else {
			logger.Info("server shutdown")
		}
	}
}

// Restart ..
func Restart() {
	// 新实例启动，等父进程(旧实例)退出后，新实例由 init 进程托管
	// 旧实例停止监听，并优雅关闭
	// gs.shutdown(gs.shutdownTimeout)
	server := gserver.server
	logger := gserver.logger

	dir, err := os.Getwd()
	if err != nil {
		logger.Fatalf("get dir failed: %s", err.Error())
	}

	// args := []string{}
	// for _, arg := range os.Args {
	// 	if arg == "-continue" {
	// 		continue
	// 	}
	// 	args = append(args, arg)
	// }
	// args = append(args, "-continue")

	files := []*os.File{os.Stdin, os.Stdout, os.Stderr}
	listenFile, err := server.ListenFile()
	if err != nil {
		logger.Fatalf("get listenFile failed: %s", err.Error())
	}
	files = append(files, listenFile)

	// ----------------- 1 syscall.ForkExec
	// execSpec := &syscall.ProcAttr{
	// 	Env:   os.Environ(),
	// 	Files: files,
	// }
	// forkID, err := syscall.ForkExec(os.Args[0], args, execSpec)
	// if err != nil {
	// 	panic(fmt.Sprintf("syscall.ForkExec failed: %s", err.Error()))
	// }
	// if gs.isLogEnable() {
	// 	gs.logger().Infof("restart success, new pid: %d", forkID)
	// }

	// ----------------- 2 os.StartProcesses
	argv0, err := exec.LookPath(os.Args[0])
	if err != nil {
		logger.Fatalf("%s look path failed: %s", os.Args[0], err.Error())
	}
	process, err := os.StartProcess(argv0, os.Args, &os.ProcAttr{
		Dir:   dir,
		Env:   os.Environ(),
		Files: files,
	})
	if err != nil {
		logger.Fatalf("start new process failed: %s", err.Error())
	}
	gserver.logger.Infof("restart success, new pid: %d", process.Pid)

	Shutdown()
}
