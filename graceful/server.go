package graceful

import (
	"net"
	"net/http"
	"os"

	"github.com/alex-my/ghelper/logger"
)

// Package 优雅的关闭 http.Server
// TODO: 优雅的重启

// Server 替代 http.Server
type Server struct {
	*http.Server
	tc tcpKeepAliveListener // 用于复制文件描述符 dup
}

// NewServer ..
func NewServer(handler http.Handler, addr string) (*Server, error) {

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	tc := tcpKeepAliveListener{ln.(*net.TCPListener)}

	// 结构体初始化: 如果匿名字段也要初始化，则采取不声明 key 的方式
	server := &Server{&http.Server{
		Addr:    addr,
		Handler: handler,
	}, tc}

	return server, nil
}

// ListenAndServe ..
func (server *Server) ListenAndServe() error {
	addr := server.Addr
	if addr == "" {
		addr = ":http"
	}

	return server.Serve(server.tc)
}

// ListenAndServeTLS ..
func (server *Server) ListenAndServeTLS(certFile, keyFile string) error {
	addr := server.Addr
	if addr == "" {
		addr = ":https"
	}

	defer server.tc.Close()

	return server.ServeTLS(server.tc, certFile, keyFile)
}

// AddServer 服务器加进来后，就可以优雅的关闭和重启
func AddServer(server *Server, opts ...Option) {
	hs := &httpServer{
		server:  server,
		timeout: -1,
	}

	// 添加额外选项
	for _, opt := range opts {
		opt(hs)
	}

	gserver.servers = append(gserver.servers, hs)
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
		server.server.RegisterOnShutdown(f)
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
