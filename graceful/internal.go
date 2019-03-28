package graceful

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/alex-my/ghelper/logger"
	"github.com/alex-my/ghelper/time"
)

// gracefulServer 管理 http.Server
type gracefulServer struct {
	servers []*http.Server

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

func (gs *gracefulServer) close() {
	if len(gs.servers) == 0 {
		return
	}

	logEnable := gs.isLogEnable()
	logger := gs.logger()

	for _, server := range gs.servers {
		err := server.Close()
		if logEnable {
			if err != nil {
				logger.Errorf("Server close: %s", err.Error())
			} else {
				logger.Infof("Server exiting")
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
		if timeout > 0 {
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second(timeout))
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				if logEnable {
					logger.Errorf("Server shutdown with timeout: %s", err.Error())
				}
			}
			select {
			case <-ctx.Done():
				if logEnable {
					logger.Infof("timeout of %d seconds", timeout)
				}
			}
		} else {
			ctx := context.TODO()
			if err := server.Shutdown(ctx); err != nil {
				if logEnable {
					logger.Errorf("Server shutdown: %s", err.Error())
				}
			}
		}
		if logEnable {
			logger.Infof("Server exiting")
		}
	}
	if logEnable {
		logger.Infof("All server exiting")
	}
}

func (gs *gracefulServer) restart() {

}

func (gs *gracefulServer) listenSignal(f func()) {
	if len(gs.servers) == 0 {
		return
	}

	if len(gs.restartSignals) == 0 {
		gs.restartSignals = []os.Signal{syscall.SIGUSR1, syscall.SIGUSR2}
	}
	if len(gs.closeSignals) == 0 {
		gs.closeSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
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
