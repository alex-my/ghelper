package graceful

import (
	"net"
	"net/http"
	"os"

	"github.com/alex-my/ghelper/time"
)

// Package 优雅的关闭 http.Server
// TODO: 优雅的重启

// Server 替代 http.Server
type Server struct {
	*http.Server
	tc      tcpKeepAliveListener // 用于复制文件描述符 dup
	name    string
	timeout int // 可以为服务器单独指定超时时间
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
	}, tc, "", -1}

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

// ListenFile 获取监听套接字文件
func (server *Server) ListenFile() (*os.File, error) {
	file, err := server.tc.File()
	if err != nil {
		return nil, nil
	}
	return file, nil
}

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}

	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(time.Minute(3))

	return tc, nil
}
