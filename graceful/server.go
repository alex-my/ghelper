package graceful

import (
	"net"
	"net/http"
	"os"
	"syscall"

	"github.com/alex-my/ghelper/logger"
	"github.com/alex-my/ghelper/time"
)

// envNewKey 新创建的进程环境变量中拥有该标签
var envNewKey = "newProcess"

// Server 替代 http.Server
type Server struct {
	*http.Server
	// tc 用于获取监控套接字文件
	tc tcpKeepAliveListener
}

// NewServer ..
func NewServer(handler http.Handler, logger logger.Log) *Server {
	var tc tcpKeepAliveListener

	// 结构体初始化: 如果匿名字段也要初始化，则采取不声明 key 的方式
	server := &Server{&http.Server{
		Handler: handler,
	}, tc}

	gserver.server = server
	gserver.logger = logger

	return server
}

// ListenAndServe ..
func (server *Server) ListenAndServe(addr string) error {
	if addr == "" {
		addr = ":http"
	}

	ln, err := listen(addr)
	if err != nil {
		return err
	}
	tc := tcpKeepAliveListener{ln.(*net.TCPListener)}

	server.Addr = addr
	server.tc = tc

	return server.Serve(server.tc)
}

// ListenAndServeTLS ..
func (server *Server) ListenAndServeTLS(addr, certFile, keyFile string) error {
	if addr == "" {
		addr = ":https"
	}

	ln, err := listen(addr)
	if err != nil {
		return err
	}
	tc := tcpKeepAliveListener{ln.(*net.TCPListener)}

	server.Addr = addr
	server.tc = tc

	defer server.tc.Close()

	return server.ServeTLS(server.tc, certFile, keyFile)
}

// listenFile 获取监听套接字文件
func (server *Server) listenFile() (*os.File, error) {
	file, err := server.tc.File()
	if err != nil {
		return nil, nil
	}
	return file, nil
}

// listen 创建监听套接字
func listen(addr string) (ln net.Listener, err error) {
	if _, ok := syscall.Getenv(envNewKey); ok {
		fp := os.NewFile(3, "")
		defer fp.Close()
		ln, err = net.FileListener(fp)
	} else {
		ln, err = net.Listen("tcp", addr)
	}
	return
}

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}

	if err = tc.SetKeepAlive(true); err != nil {
		return nil, err
	}
	if err = tc.SetKeepAlivePeriod(time.Minute(3)); err != nil {
		return nil, err
	}

	return tc, nil
}
